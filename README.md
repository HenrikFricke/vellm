# Vellm
> CLI tool for Medium

## Install

Install vellm easily with your Go environment:

```bash
go install github.com/HenrikFricke/vellm
```

## Usage

Vellm needs an integration token to work with Medium. Run `vellm setup` for more information 
and to add the token. After you setup Vellm you can upload Markdown files in seconds.

1. Create a Markdown file: `touch mystory.md`
1. Write your story, for example:
    ```md
    ---
    Title: This is a title
    Tags:
        - go
        - lang
        - development
    ---
    # This is a title

    Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis nec consectetur purus. Integer malesuada nunc vitae nisl porttitor, id aliquet nibh facilisis. 
    ```
2. Run `vellm publish -f ./mystory.md`
3. Check your new story on Medium and publish it :)

You can define several meta data with a [Front Matter]. Vellm accepts the following options:

| Key | Format | Example | Default |
|---|---|---|---|
| Title | string | `Title: This is a title` | - |
| Tags | []string | <code>Tags:</code><br /><code>&nbsp;&nbsp;&nbsp;&nbsp;- golang</code><br /><code>&nbsp;&nbsp;&nbsp;&nbsp;- development</code> | - |
| Published | boolean | `Published: true` | false |

[Front Matter]: http://assemble.io/docs/YAML-front-matter.html