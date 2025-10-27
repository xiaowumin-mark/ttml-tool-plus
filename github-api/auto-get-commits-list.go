package githubapi

import "log"

func (g *GithubApiService) AutoGetCommitsList(now, to string) ([]*CommitItem, []*CommitItemFile, error) {
	nextSha := to
	var commits []*CommitItem
	for nextSha != now {
		commit, err := g.GetCommit(nextSha)
		if err != nil {
			return nil, nil, err
		}
		nextSha = commit.Parents[0].Sha
		commits = append(commits, commit)
		log.Println(commit.Sha)
	}
	// 倒序commits
	for i, j := 0, len(commits)-1; i < j; i, j = i+1, j-1 {
		commits[i], commits[j] = commits[j], commits[i]
	}
	var files []*CommitItemFile
	for _, commit := range commits {
		for _, file := range commit.Files {
			iss, _ := fileIsContains(files, file.FileName)
			if iss {
				// 文件已存在 则更新
				filed := getFile(files, file.FileName)
				filed.Sha = file.Sha
				filed.FileName = file.FileName
				filed.AddiTions = file.AddiTions
				filed.Deletions = file.Deletions
				filed.Changes = file.Changes
				filed.BlobURL = file.BlobURL
				filed.RawURL = file.RawURL
				filed.ContentsURL = file.ContentsURL
				filed.Patch = file.Patch
			} else {
				files = append(files, &file)
			}
		}
	}
	return commits, files, nil

}

func fileIsContains(files []*CommitItemFile, fileName string) (bool, *CommitItemFile) {
	for _, file := range files {
		if file.FileName == fileName {
			return true, file
		}
	}
	return false, nil
}

func getFile(files []*CommitItemFile, fileName string) *CommitItemFile {
	for _, file := range files {
		if file.FileName == fileName {
			return file
		}
	}
	return nil
}
