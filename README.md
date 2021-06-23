# DingTalk Github action

This github action sends a dingtalk message.

## Example usage

```
uses: hugozhu/dingtalk-github-action@v1
with:
  corpid: ${{ secrets.CORP_ID }}
  corpsecret: ${{ secrets.CORP_SECRET }}
  token: ${{ secrets.TOKEN }}
  title: 'Hello from Github action'
  file: 'test.output'
```
