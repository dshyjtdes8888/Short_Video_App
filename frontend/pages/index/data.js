//1.{}里包裹的是每一个视频的参数
const userList = [{//1
	"_id": "60ee559d9dad8509164E5",//1.每一个视频独有 id （自定义）
	"username": "井川里予",//2.视频拥有者名称
	"href": "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fup.enterdesk.com%2Fedpic%2F39%2Fb7%2F53%2F39b75357f98675e2d6d5dcde1fb805a3.jpg&refer=http%3A%2F%2Fup.enterdesk.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1642766597&t=a3d7a257c1a4666088a09dc517a1473a",//3.头像
	"title": "mood",//4.第一行标题
	"msg": "#vlog日常 #抖音热门 #甜系女孩#圣诞摇铃变装 后续米了哟!提刖拍主诞卷死你们",//5.第二行内容
	"state": "pause",//6.初始状态标志（不改）
	"like": 0,//7.是否是自己喜欢，0不喜欢，1喜欢。默认不喜欢
	"like_n": 0,//8.喜欢数量
	"sms_n": 0,//9.评论数
	"src": "https://cg-test-video.oss-cn-hangzhou.aliyuncs.com/douyin/%E5%81%87%E5%A6%82%E4%BD%A0%E7%9A%84%E5%89%8D%E5%A5%B3%E5%8F%8B%E5%92%8C%E7%8E%B0%E5%A5%B3%E5%8F%8B%E5%90%8C%E6%97%B6%E6%8E%89%E8%BF%9B%E4%BA%86%E6%88%91%E8%BA%AB%E5%90%8E%E7%9A%84%E8%BF%99%E7%89%87%E6%B5%B7%EF%BC%8C%E8%AF%B7%E9%97%AE%E2%80%A6%E6%88%91%E8%83%BD%E8%B7%9F%E4%BD%A0%E5%BC%80%E4%B8%AA%E9%BB%91%E5%90%97%EF%BC%9F.mp4",//10.视频链接
	"playNumber": 0,//11.播放视频的数量
	"pinlun": [],//12.评论
  "dubbedText":'@ET外星人街舞创作的原声-ET',//配音文字
	"playIng": false,//13.播放（默认这个即可）
	"isShowimage": false,//14.是否显示封面（默认这个即可）
	"isShowProgressBarTime": false,//15.是否显示进度条（默认这个即可）
	"isplay": true//16.是否播放音频（默认这个即可）
},{//1
	"_id": "60ee559dhde0ad8509164E5",//1.每一个视频独有 id （自定义）
	"username": "遇见",//2.视频拥有者名称
	"href": "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fup.enterdesk.com%2Fedpic%2F39%2Fb7%2F53%2F39b75357f98675e2d6d5dcde1fb805a3.jpg&refer=http%3A%2F%2Fup.enterdesk.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1642766597&t=a3d7a257c1a4666088a09dc517a1473a",//3.头像
	"title": "mood",//4.第一行标题
	"msg": "moodmoodmoodmoodmoodmoodmoodmoodmoodmoodmoodmoodmoodmoodmoodmoodmoodmood",//5.第二行内容
	"state": "pause",//6.初始状态标志（不改）
	"like": 0,//7.是否是自己喜欢，0不喜欢，1喜欢。默认不喜欢
	"like_n": 0,//8.喜欢数量
	"sms_n": 0,//9.评论数
	"src": "https://cg-test-video.oss-cn-hangzhou.aliyuncs.com/douyin/%E6%8B%96%E9%9E%8B%E5%BD%92%E5%B1%9E%E6%84%9F.mp4",//10.视频链接
	"playNumber": 0,//11.播放视频的数量
	"pinlun": [],//12.评论
  "dubbedText":'@遇见创作的原声 - 遇见',//配音文字
	"playIng": false,//13.播放（默认这个即可）
	"isShowimage": false,//14.是否显示封面（默认这个即可）
	"isShowProgressBarTime": false,//15.是否显示进度条（默认这个即可）
	"isplay": true//16.是否播放音频（默认这个即可）
},{//1
	"_id": "60ee5014faad8509164E5",//1.每一个视频独有 id （自定义）
	"username": "意识到",//2.视频拥有者名称
	"href": "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fup.enterdesk.com%2Fedpic%2F39%2Fb7%2F53%2F39b75357f98675e2d6d5dcde1fb805a3.jpg&refer=http%3A%2F%2Fup.enterdesk.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1642766597&t=a3d7a257c1a4666088a09dc517a1473a",//3.头像
	"title": "mood",//4.第一行标题
	"msg": "moodmoodmoodmoodmoodmoodmoodmoodmoodmoodmoodmoodmoodmoodmoodmoodmoodmood",//5.第二行内容
	"state": "pause",//6.初始状态标志（不改）
	"like": 0,//7.是否是自己喜欢，0不喜欢，1喜欢。默认不喜欢
	"like_n": 0,//8.喜欢数量
	"sms_n": 0,//9.评论数
	"src": "https://cg-test-video.oss-cn-hangzhou.aliyuncs.com/douyin/%E8%BF%99%E4%B8%80%E5%B9%B4%E6%9C%89%E4%BB%80%E4%B9%88%E9%9A%BE%E5%BF%98%E7%9A%84%E5%9B%9E%E5%BF%86%E4%B9%88.mp4",//10.视频链接
	"playNumber": 0,//11.播放视频的数量
	"pinlun": [],//12.评论
  "dubbedText":'@遇见创作的原声 - 遇见',//配音文字
	"playIng": false,//13.播放（默认这个即可）
	"isShowimage": false,//14.是否显示封面（默认这个即可）
	"isShowProgressBarTime": false,//15.是否显示进度条（默认这个即可）
	"isplay": true//16.是否播放音频（默认这个即可）
},{//1
	"_id": "12ee559d9dad8509164E5",//1.每一个视频独有 id （自定义）
	"username": "意识到",//2.视频拥有者名称
	"href": "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fup.enterdesk.com%2Fedpic%2F39%2Fb7%2F53%2F39b75357f98675e2d6d5dcde1fb805a3.jpg&refer=http%3A%2F%2Fup.enterdesk.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1642766597&t=a3d7a257c1a4666088a09dc517a1473a",//3.头像
	"title": "mood",//4.第一行标题
	"msg": "moodmoodmoodmoodmoodmoodmoodmoodmoodmoodmoodmoodmoodmoodmoodmoodmoodmood",//5.第二行内容
	"state": "pause",//6.初始状态标志（不改）
	"like": 0,//7.是否是自己喜欢，0不喜欢，1喜欢。默认不喜欢
	"like_n": 0,//8.喜欢数量
	"sms_n": 0,//9.评论数
	"src": "https://cg-test-video.oss-cn-hangzhou.aliyuncs.com/douyin/%E5%90%83%E9%B8%AD%E8%85%BF.mp4",//10.视频链接
	"playNumber": 0,//11.播放视频的数量
	"pinlun": [],//12.评论
  "dubbedText":'@遇见创作的原声 - 遇见',//配音文字
	"playIng": false,//13.播放（默认这个即可）
	"isShowimage": false,//14.是否显示封面（默认这个即可）
	"isShowProgressBarTime": false,//15.是否显示进度条（默认这个即可）
	"isplay": true//16.是否播放音频（默认这个即可）
},{//1
	"_id": "60hj5dsdad8509164E5",//1.每一个视频独有 id （自定义）
	"username": "意识到",//2.视频拥有者名称
	"href": "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fup.enterdesk.com%2Fedpic%2F39%2Fb7%2F53%2F39b75357f98675e2d6d5dcde1fb805a3.jpg&refer=http%3A%2F%2Fup.enterdesk.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1642766597&t=a3d7a257c1a4666088a09dc517a1473a",//3.头像
	"title": "mood",//4.第一行标题
	"msg": "moodmoodmoodmoodmoodmoodmoodmoodmoodmoodmoodmoodmoodmoodmoodmoodmoodmood",//5.第二行内容
	"state": "pause",//6.初始状态标志（不改）
	"like": 0,//7.是否是自己喜欢，0不喜欢，1喜欢。默认不喜欢
	"like_n": 0,//8.喜欢数量
	"sms_n": 0,//9.评论数
	"src": "https://cg-test-video.oss-cn-hangzhou.aliyuncs.com/douyin/%E4%BD%A0%E4%BB%A5%E4%B8%BA%E4%B8%BB%E6%92%AD.mp4",//10.视频链接
	"playNumber": 0,//11.播放视频的数量
	"pinlun": [],//12.评论
  "dubbedText":'@遇见创作的原声 - 遇见',//配音文字
	"playIng": false,//13.播放（默认这个即可）
	"isShowimage": false,//14.是否显示封面（默认这个即可）
	"isShowProgressBarTime": false,//15.是否显示进度条（默认这个即可）
	"isplay": true//16.是否播放音频（默认这个即可）
},{//1
	"_id": "60ee559d90gd8509164E5",//1.每一个视频独有 id （自定义）
	"username": "意识到",//2.视频拥有者名称
	"href": "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fup.enterdesk.com%2Fedpic%2F39%2Fb7%2F53%2F39b75357f98675e2d6d5dcde1fb805a3.jpg&refer=http%3A%2F%2Fup.enterdesk.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1642766597&t=a3d7a257c1a4666088a09dc517a1473a",//3.头像
	"title": "mood",//4.第一行标题
	"msg": "moodmoodmoodmoodmoodmoodmoodmoodmoodmoodmoodmoodmoodmoodmoodmoodmoodmood",//5.第二行内容
	"state": "pause",//6.初始状态标志（不改）
	"like": 0,//7.是否是自己喜欢，0不喜欢，1喜欢。默认不喜欢
	"like_n": 0,//8.喜欢数量
	"sms_n": 0,//9.评论数
	"src": "https://cg-test-video.oss-cn-hangzhou.aliyuncs.com/douyin/%E7%BD%97%E6%B0%B8%E6%B5%A9.mp4",//10.视频链接
	"playNumber": 0,//11.播放视频的数量
	"pinlun": [],//12.评论
  "dubbedText":'@遇见创作的原声 - 遇见',//配音文字
	"playIng": false,//13.播放（默认这个即可）
	"isShowimage": false,//14.是否显示封面（默认这个即可）
	"isShowProgressBarTime": false,//15.是否显示进度条（默认这个即可）
	"isplay": true//16.是否播放音频（默认这个即可）
},{//3
	"_id": "60edc229f3ff4f000138cb76",
	"username": "阿斯达达",
	"href": "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fup.enterdesk.com%2Fedpic%2F39%2Fb7%2F53%2F39b75357f98675e2d6d5dcde1fb805a3.jpg&refer=http%3A%2F%2Fup.enterdesk.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1642766597&t=a3d7a257c1a4666088a09dc517a1473a",//3.头像
	"title": "",
	"msg": "最向往的地方-林俊杰",
	"state": "pause",
	"like": 0,
	"like_n": 0,
	"sms_n": 0,
	"src": "https://vkceyugu.cdn.bspapp.com/VKCEYUGU-0455454d-b373-4768-aa39-dc1226fc1362/53543262-55f5-4685-a5e3-b56ce75bcb88.mp4",
	"playNumber": 0,
	"pinlun": [],
  "dubbedText":'@阿斯达达的原声 - 阿斯达达',//配音文字
	"playIng": false,
	"isShowimage": false,
	"isShowProgressBarTime": false,
	"isplay": true
}, {//4
	"_id": "60edc16570f7740001aa165a",
	"username": "阿斯顿",
	"href": "https://vkceyugu.cdn.bspapp.com/VKCEYUGU-2908110e-6da2-4899-8b44-d45c153457ad/247c9ea8-c708-440b-9dbf-3c3facb728ee.png",
	"title": "",
	"msg": "Ring Ring Ring",
	"state": "pause",
	"like": 0,
	"like_n": 0,
	"sms_n": 0,
	"src": "https://vkceyugu.cdn.bspapp.com/VKCEYUGU-0455454d-b373-4768-aa39-dc1226fc1362/bfc86ab8-bb3b-4cef-a5d2-8c5edce4ef17.mp4",
	"playNumber": 0,
	"pinlun": [],
  "dubbedText":'@阿斯顿的原声 - 阿斯顿',//配音文字
	"playIng": false,
	"isShowimage": false,
	"isShowProgressBarTime": false,
	"isplay": true
}, {//5
	"_id": "60edbacf5e31f50001a66f9c",
	"username": "史蒂夫",
	"href": "https://vkceyugu.cdn.bspapp.com/VKCEYUGU-2908110e-6da2-4899-8b44-d45c153457ad/247c9ea8-c708-440b-9dbf-3c3facb728ee.png",
	"title": "",
	"msg": "有何不可",
	"state": "pause",
	"like": 0,
	"like_n": 0,
	"sms_n": 3,
	"src": "https://vkceyugu.cdn.bspapp.com/VKCEYUGU-0455454d-b373-4768-aa39-dc1226fc1362/5017a17a-389b-45e0-8d91-711c9dc76759.mp4",
	"playNumber": 0,
	"pinlun": [],
  "dubbedText":'@史蒂夫的原声 - 史蒂夫',//配音文字
	"playIng": false,
	"isShowimage": false,
	"isShowProgressBarTime": false,
	"isplay": true
}, {//6
	"_id": "60ed95fa01033700016f1acf",
	"username": "很久很久",
	"href": "https://vkceyugu.cdn.bspapp.com/VKCEYUGU-2908110e-6da2-4899-8b44-d45c153457ad/247c9ea8-c708-440b-9dbf-3c3facb728ee.png",
	"title": "",
	"msg": "破茧",
	"state": "pause",
	"like": 0,
	"like_n": 0,
	"sms_n": 2,
	"src": "https://vkceyugu.cdn.bspapp.com/VKCEYUGU-0455454d-b373-4768-aa39-dc1226fc1362/209180d8-3dfd-42ea-9ef5-5f98ae0d95e1.mp4",
	"playNumber": 0,
	"pinlun": [],
  "dubbedText":'@很久很久的原声 - 很久很久',//配音文字
	"playIng": false,
	"isShowimage": false,
	"isShowProgressBarTime": false,
	"isplay": true
}, {//7
	"_id": "60e8dd308a69dc0001f1435a",
	"username": "汪苏泷",
	"href": "https://vkceyugu.cdn.bspapp.com/VKCEYUGU-2908110e-6da2-4899-8b44-d45c153457ad/247c9ea8-c708-440b-9dbf-3c3facb728ee.png",
	"title": "",
	"msg": "后会无期-徐良-汪苏泷",
	"state": "pause",
	"like": 0,
	"like_n": 0,
	"sms_n": 1,
	"src": "https://vkceyugu.cdn.bspapp.com/VKCEYUGU-0455454d-b373-4768-aa39-dc1226fc1362/c8f7a17f-6eb8-453a-9f43-944ecc7a9f11.mp4",
	"playNumber": 0,
	"pinlun": [],
  "dubbedText":'@汪苏泷的原声 - 汪苏泷',//配音文字
	"playIng": false,
	"isShowimage": false,
	"isShowProgressBarTime": false,
	"isplay": true
}, {//8
	"_id": "60e86d052881d300012eb3e9",
	"username": "哈哈哈哈哈",
	"href": "https://vkceyugu.cdn.bspapp.com/VKCEYUGU-2908110e-6da2-4899-8b44-d45c153457ad/247c9ea8-c708-440b-9dbf-3c3facb728ee.png",
	"title": "",
	"msg": "笑死了哈哈哈哈哈",
	"state": "pause",
	"like": 0,
	"like_n": 0,
	"sms_n": 0,
	"src": "https://vkceyugu.cdn.bspapp.com/VKCEYUGU-2908110e-6da2-4899-8b44-d45c153457ad/f905c750-c70e-46b2-aaa6-37778d308f13.mp4",
	"playNumber": 0,
	"pinlun": [],
  "dubbedText":'@哈哈哈哈哈的原声 - 哈哈哈哈哈',//配音文字
	"playIng": false,
	"isShowimage": false,
	"isShowProgressBarTime": false,
	"isplay": true
}, {//9
	"_id": "60e86b658a69dc0001f003ae",
	"username": "望·即",
	"href": "https://vkceyugu.cdn.bspapp.com/VKCEYUGU-2908110e-6da2-4899-8b44-d45c153457ad/247c9ea8-c708-440b-9dbf-3c3facb728ee.png",
	"title": "",
	"msg": "Justin Biber- As long as you love me",
	"state": "pause",
	"like": 0,
	"like_n": 0,
	"sms_n": 0,
	"src": "https://vkceyugu.cdn.bspapp.com/VKCEYUGU-0455454d-b373-4768-aa39-dc1226fc1362/9392e85c-36db-473f-8ec3-4f8ed83a382a.mp4",
	"playNumber": 0,
	"pinlun": [],
  "dubbedText":'@望·即的原声 - 望·即',//配音文字
	"playIng": false,
	"isShowimage": false,
	"isShowProgressBarTime": false,
	"isplay": true
}, {//10
	"_id": "60e86a9e9056cd000110da51",
	"username": "user",
	"href": "https://vkceyugu.cdn.bspapp.com/VKCEYUGU-2908110e-6da2-4899-8b44-d45c153457ad/247c9ea8-c708-440b-9dbf-3c3facb728ee.png",
	"title": "",
	"msg": "苹果2019年发布会开场动画",
	"state": "pause",
	"like": 0,
	"like_n": 0,
	"sms_n": 0,
	"src": "https://vkceyugu.cdn.bspapp.com/VKCEYUGU-0455454d-b373-4768-aa39-dc1226fc1362/e1cd785e-56ae-4c96-a713-126bf2950e19.mp4",
	"playNumber": 0,
	"pinlun": [],
  "dubbedText":'@阿斯达达的原声 - 阿斯达达',//配音文字
	"playIng": false,
	"isShowimage": false,
	"isShowProgressBarTime": false,
	"isplay": true
}, {
	//11
	"_id": "60e86a9e9056cd000110da5q",
	"username": "user",
	"href": "https://vkceyugu.cdn.bspapp.com/VKCEYUGU-2908110e-6da2-4899-8b44-d45c153457ad/247c9ea8-c708-440b-9dbf-3c3facb728ee.png",
	"title": "",
	"msg": "每天笑一笑",
	"state": "pause",
	"like": 0,
	"like_n": 0,
	"sms_n": 0,
	"src": "https://vkceyugu.cdn.bspapp.com/VKCEYUGU-2908110e-6da2-4899-8b44-d45c153457ad/97b50a6d-f77d-4418-b38d-844e0b9eec97.mp4",
	"playNumber": 0,
	"pinlun": [],
  "dubbedText":'@阿斯达达的原声 - 阿斯达达',//配音文字
	"playIng": false,
	"isShowimage": false,
	"isShowProgressBarTime": false,
	"isplay": true
}, {
	//12
	"_id": "60e86a9e9026cd000119da5q",
	"username": "user",
	"href": "https://vkceyugu.cdn.bspapp.com/VKCEYUGU-2908110e-6da2-4899-8b44-d45c153457ad/247c9ea8-c708-440b-9dbf-3c3facb728ee.png",
	"title": "",
	"msg": "老外真会玩",
	"state": "pause",
	"like": 0,
	"like_n": 0,
	"sms_n": 0,
	"src": "https://vkceyugu.cdn.bspapp.com/VKCEYUGU-2908110e-6da2-4899-8b44-d45c153457ad/c061b07a-4b34-4d6d-aa1a-2cf41679f17c.mp4",
	"playNumber": 0,
	"pinlun": [],
  "dubbedText":'@阿斯达达的原声 - 阿斯达达',//配音文字
	"playIng": false,
	"isShowimage": false,
	"isShowProgressBarTime": false,
	"isplay": true
}]
export default userList;