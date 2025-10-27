package gitapi

import (
	"fmt"
	"log"
	"testing"

	"github.com/go-git/go-git/v6"
)

func TestGoGit(t *testing.T) {
	//repoPath := `E:\projects\test\amll-ttml-db`
	repoPath := `C:\Users\xiaow\.ttml-tool-plus\ttml-db`
	// 打开已有仓库，不存在就克隆
	repo, err := git.PlainOpen(repoPath)
	if err == git.ErrRepositoryNotExists {
		/*fmt.Println("正在克隆仓库...")
		repo, err = git.PlainClone(repoPath, false, &git.CloneOptions{
			URL:      "https://github.com/go-git/go-git",
			Progress: nil,
		})
		if err != nil {
			log.Fatal(err)
		}*/

		log.Println("没有仓库")
		return
	}

	// 先 fetch 拉取远程最新
	fmt.Println("Fetching...")
	err = repo.Fetch(&git.FetchOptions{
		RemoteName: "origin",
		Force:      true,
	})
	if err != nil && err != git.NoErrAlreadyUpToDate {
		log.Fatal(err)
	}
	// 获取工作区状态
	w, _ := repo.Worktree()
	status, _ := w.Status()

	fmt.Println("=== 工作区文件变化 ===")
	for file, s := range status {
		switch {
		case s.Staging == git.Modified || s.Worktree == git.Modified:
			fmt.Println("修改:", file)
		case s.Staging == git.Added || s.Worktree == git.Added || s.Staging == git.Renamed:
			fmt.Println("新增:", file)
		case s.Staging == git.Deleted || s.Worktree == git.Deleted:
			fmt.Println("删除:", file)
		case s.Staging == git.Untracked:
			fmt.Println("未跟踪:", file)
		}
	}

	// 对比本地与远程差异
	/*localRef, _ := repo.Head()
	remoteRef, _ := repo.Reference(
		plumbing.NewRemoteReferenceName("origin", localRef.Name().Short()),
		true,
	)

	if localRef.Hash() != remoteRef.Hash() {
		fmt.Println("\n=== 本地 vs 远程差异 ===")
		remoteCommit, _ := repo.CommitObject(remoteRef.Hash())
		localCommit, _ := repo.CommitObject(localRef.Hash())
		patch, _ := localCommit.Patch(remoteCommit)
		fmt.Println(patch)
	} else {
		fmt.Println("\n本地与远程一致 ✅")
	}*/
}
