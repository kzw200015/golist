# golist
一个简单的列目录程序，Go实现，低占用
支持文本预览、代码高亮、音频预览、视频预览
## 使用
```
golist -p 8080 -s "/data"
```
然后打开`http://127.0.0.1:8080`
## 构建
### 拉取源码
```
git clone --recurse-submodules https://github.com/kzw200015/golist.git
```
### 构建前端
```
cd golist/assets/frontend
yarn && yarn build
```
### 打包构建前后端
```
cd golist
go build -o golist
```