package disk

import (
	"fmt"
	"github.com/owncloud/ocis/accounts/pkg/config"
	"github.com/owncloud/ocis/accounts/pkg/indexer/errors"
	"github.com/owncloud/ocis/accounts/pkg/indexer/index"
	"github.com/owncloud/ocis/accounts/pkg/indexer/option"
	. "github.com/owncloud/ocis/accounts/pkg/indexer/test"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"
)

func TestNonUniqueIndexAdd(t *testing.T) {
	sut, dataPath := getNonUniqueIdxSut(t, TestPet{}, "Color")

	ids, err := sut.Lookup("Green")
	assert.NoError(t, err)
	assert.EqualValues(t, []string{"goefe-789", "xadaf-189"}, ids)

	ids, err = sut.Lookup("White")
	assert.NoError(t, err)
	assert.EqualValues(t, []string{"wefwe-456"}, ids)

	ids, err = sut.Lookup("Cyan")
	assert.Error(t, err)
	assert.EqualValues(t, []string{}, ids)

	_ = os.RemoveAll(dataPath)

}

func TestNonUniqueIndexUpdate(t *testing.T) {
	sut, dataPath := getNonUniqueIdxSut(t, TestPet{}, "Color")

	err := sut.Update("goefe-789", "Green", "Black")
	assert.NoError(t, err)

	err = sut.Update("xadaf-189", "Green", "Black")
	assert.NoError(t, err)

	assert.DirExists(t, path.Join(dataPath, fmt.Sprintf("index.disk/non_unique.%v.Color/Black", GetTypeFQN(TestPet{}))))
	assert.NoDirExists(t, path.Join(dataPath, fmt.Sprintf("index.disk/non_unique.%v.Color/Green", GetTypeFQN(TestPet{}))))

	_ = os.RemoveAll(dataPath)
}

func TestNonUniqueIndexDelete(t *testing.T) {
	sut, dataPath := getNonUniqueIdxSut(t, TestPet{}, "Color")
	assert.FileExists(t, path.Join(dataPath, fmt.Sprintf("index.disk/non_unique.%v.Color/Green/goefe-789", GetTypeFQN(TestPet{}))))

	err := sut.Remove("goefe-789", "")
	assert.NoError(t, err)
	assert.NoFileExists(t, path.Join(dataPath, fmt.Sprintf("index.disk/non_unique.%v.Color/Green/goefe-789", GetTypeFQN(TestPet{}))))

	_ = os.RemoveAll(dataPath)
}

func TestNonUniqueIndexSearch(t *testing.T) {
	sut, dataPath := getNonUniqueIdxSut(t, TestPet{}, "Email")

	res, err := sut.Search("Gr*")

	assert.NoError(t, err)
	assert.Len(t, res, 2)

	assert.Equal(t, "goefe-789", path.Base(res[0]))
	assert.Equal(t, "xadaf-189", path.Base(res[1]))

	res, err = sut.Search("does-not-exist@example.com")
	assert.Error(t, err)
	assert.IsType(t, &errors.NotFoundErr{}, err)

	_ = os.RemoveAll(dataPath)
}

// entity: used to get the fully qualified name for the index root path.
func getNonUniqueIdxSut(t *testing.T, entity interface{}, indexBy string) (index.Index, string) {
	dataPath := WriteIndexTestData(t, TestData, "Id")
	cfg := config.Config{
		Repo: config.Repo{
			Disk: config.Disk{
				Path: dataPath,
			},
		},
	}

	sut := NewNonUniqueIndexWithOptions(
		option.WithTypeName(GetTypeFQN(entity)),
		option.WithIndexBy(indexBy),
		option.WithFilesDir(path.Join(cfg.Repo.Disk.Path, "pets")),
		option.WithDataDir(cfg.Repo.Disk.Path),
	)
	err := sut.Init()
	if err != nil {
		t.Fatal(err)
	}

	for _, u := range TestData["pets"] {
		pkVal := ValueOf(u, "Id")
		idxByVal := ValueOf(u, "Color")
		_, err := sut.Add(pkVal, idxByVal)
		if err != nil {
			t.Fatal(err)
		}
	}

	return sut, dataPath
}
