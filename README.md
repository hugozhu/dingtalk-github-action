# DingTalk Github action

This github action sends a dingtalk message.

## Example usage

```
uses: hugozhu/dingtalk-github-action@v1
with:
  token: ${{ secrets.TOKEN }}
  title: 'Hello from Github action'
  type: markdown
  file: 'test.output'
```


```
uses: hugozhu/dingtalk-github-action@v1
with:
  token: ${{ secrets.TOKEN }}
  title: 'Hello from Github action'
  type: text
  content: |
    ## Hello
    > from github action message
```