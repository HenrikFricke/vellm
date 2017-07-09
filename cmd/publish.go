package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"

	medium "github.com/Medium/medium-sdk-go"
	"github.com/ericaro/frontmatter"
	"github.com/spf13/cobra"
)

var markdownFilePath string

type markdownFrontmatter struct {
	Title     string   `yaml:"Title"`
	Tags      []string `yaml:"Tags"`
	Published bool     `yaml:"Published"`
	Content   string   `fm:"content" yaml:"-"`
}

// publishCmd represents the publish command
var publishCmd = &cobra.Command{
	Use:     "publish",
	Short:   "Publish a Markdown file to Medium.",
	Long:    `Publish a Markdown file to Medium.`,
	PreRunE: mediumTokenCheck,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error

		if markdownFilePath == "" {
			return errors.New("Path to Markdown file is missing.")
		}

		var fileConent []byte
		if fileConent, err = ioutil.ReadFile(markdownFilePath); err != nil {
			return err
		}

		content := new(markdownFrontmatter)
		if err = frontmatter.Unmarshal(fileConent, content); err != nil {
			return err
		}

		var u *medium.User
		if u, err = mediumClient.GetUser(""); err != nil {
			return err
		}

		post, err := mediumClient.CreatePost(medium.CreatePostOptions{
			UserID:        u.ID,
			Title:         content.Title,
			Tags:          content.Tags,
			Content:       content.Content,
			ContentFormat: medium.ContentFormatMarkdown,
			PublishStatus: getPublishStatus(content.Published)})

		if err != nil {
			return err
		}

		fmt.Println("Your story has been successfully uploaded to Medium. You can find it here:")
		fmt.Println(post.URL)

		return nil
	},
}

func init() {
	RootCmd.AddCommand(publishCmd)

	publishCmd.Flags().StringVarP(&markdownFilePath, "file", "f", "", "Path to Markdown file")
}

func getPublishStatus(isPublished bool) medium.PublishStatus {
	if isPublished {
		return medium.PublishStatusPublic
	}

	return medium.PublishStatusDraft
}
