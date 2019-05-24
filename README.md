# 周报

## 投稿方式

### GitHub Issue

前往[带有 working 标签的 issue](https://github.com/dyweb/weekly/labels/working) 中，分享你看到的有趣/有价值的项目/博客/论文/新闻等

### Bookmarklet

Add the following target to your bookmark bar.

Remember to replace `%24YOUR_TOKEN_HERE` with your Github personal access token. The token could be acquired at `Settings -> Developer Settings -> Personal Access Token` and check `public_repo` permission.

```js
javascript:(function()%7Bconst%20apiRequest%20%3D%20(src%2C%20param)%20%3D%3E%20fetch('https%3A%2F%2Fapi.github.com%2F'%20%2B%20src%2C%20Object.assign(%7Bheaders%3A%20%7B%20Accept%3A%20'application%2Fvnd.github.v3%2Bjson'%2C%20'Content-Type'%3A%20'application%2Fjson'%2C%20Authorization%3A%20'token%20%24YOUR_TOKEN_HERE'%20%7D%7D%2C%20param))%3BapiRequest('repos%2Fdyweb%2Fweekly%2Fissues').then(v%20%3D%3E%20v.json()).then(d%20%3D%3E%20d.filter(item%20%3D%3E%20item.labels.some(lab%20%3D%3E%20lab.name%20%3D%3D%3D%20'working'))%5B0%5D.number).then(issuenum%20%3D%3E%20%7Bconst%20url%20%3D%20'repos%2Fdyweb%2Fweekly%2Fissues%2F'%20%2B%20issuenum%20%2B%20'%2Fcomments'%3Bconst%20description%20%3D%20window.prompt('Please%20input%20your%20description%20of%20this%20web%20page%3A'%2C%20'')%3Bconst%20body%20%3D%20description%20%2B%20'%5Cn%5Cn'%20%2B%20window.location.href%20%2B%20'%5Cn%5Cn%20*Submitted%20via%20%5Bbookmarklet%5D(https%3A%2F%2Fgist.github.com%2Fhtfy96%2F301ae2b1c477a4a644e943bbc27c9588)*%20%3Asparkles%3A'%3Breturn%20apiRequest(url%2C%20%7Bmethod%3A%20'POST'%2Cbody%3A%20JSON.stringify(%7Bbody%7D)%7D)%7D).then(()%20%3D%3E%20window.alert('Submission%20OK!')).catch(ex%20%3D%3E%20%7Bconsole.error(ex)%3Bwindow.alert('Failed%20to%20submit...%20See%20console%20log%20for%20exception')%3B%7D)%7D)()
```

## Source code
```js
const apiRequest = (src, param) => fetch('https://api.github.com/' + src, Object.assign({headers: { Accept: 'application/vnd.github.v3+json', 'Content-Type': 'application/json', Authorization: 'token $YOUR_TOKEN_HERE' }}, param));

apiRequest('repos/dyweb/weekly/issues')
    .then(v => v.json())
    .then(d => d.filter(item => item.labels.some(lab => lab.name === 'working'))[0].number) // current issue number
    .then(issuenum => {
        const url = 'repos/dyweb/weekly/issues/' + issuenum + '/comments';
        const description = window.prompt('Please input your description of this web page:', '');
        const body = description + '\n\n' + window.location.href + '\n\n *Submitted via [bookmarklet](https://gist.github.com/htfy96/301ae2b1c477a4a644e943bbc27c9588)* :sparkles:';
        return apiRequest(url, {
            method: 'POST',
            body: JSON.stringify({body})
        })
        })
    .then(() => window.alert('Submission OK!'))
    .catch(ex => {
        console.error(ex);
        window.alert('Failed to submit... See console log for exception');
     })
```

Copied from [Simple bookmarklet to send current webpage to weekly](https://gist.github.com/htfy96/301ae2b1c477a4a644e943bbc27c9588)

:tada: Special Thanks to [@htfy96][]!

## 周报维护方法

如果你是周报的维护者，请参考[周报维护方法文档](maintenance.md)了解如何维护周报系统。
