---
sidebar_position: 6
---

# Github

The [**Github Retriever**](https://pkg.go.dev/github.com/juliandanciu/go-feature-flag/retriever/githubretriever/#Retriever)
will perform an HTTP Request with your GitHub configuration to get your flags.

!!! Tip
    GitHub has rate limits, be sure to correctly set your `PollingInterval` to avoid reaching the limit.

## Example

```go linenums="1"
err := ffclient.Init(ffclient.Config{
    PollingInterval: 3 * time.Second,
    Retriever: &githubretriever.Retriever{
        RepositorySlug: "thomaspoignant/go-feature-flag",
        Branch: "main",
        FilePath: "testdata/flag-config.yaml",
        GithubToken: "XXXX",
        Timeout: 2 * time.Second,
    },
})
defer ffclient.Close()
```

## Configuration fields

To configure the access to your GitHub file:

| Field                | Description                                                                                                                                                                                                                                               |
|----------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| **`RepositorySlug`** | Your GitHub slug `org/repo-name`.                                                                                                                                                                                                                         |
| **`FilePath`**       | The path of your file.                                                                                                                                                                                                                                    |
| **`Branch`**         | *(optional)*<br/>The branch where your file is.<br/>Default: `main`                                                                                                                                                                                         |
| **`GithubToken`**    | *(optional)*<br/>Github token is used to access a private repository, you need the `repo` permission *([how to create a GitHub token](https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/creating-a-personal-access-token))*. |
| **`Timeout`**        | *(optional)*<br/>Timeout for the HTTP call <br/>Default: 10 seconds                                                                                                                                                                                         |
