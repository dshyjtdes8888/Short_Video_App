<template>
  <div style="letter-spacing: 1px">
	  <!-- 评论列表容器，设置固定高度，使得评论列表可以在容器内滚动 -->
	      <div class="comments-container">
    <div class="content">
      <!-- 一级评论 -->
      <div class="first" v-for="(item, index) in comments" :key="index">
        <a href="javascript:;" class="first-img">
          <!-- 如果当前用户有头像则显示头像,没有则显示默认头像 -->
          <img v-if="item.avatar" :src="item.avatar" alt="" />
          <img v-else src="https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg" alt="" />
        </a>
        <div class="first-content">
          <!-- 一级评论用户昵称 -->
          <h3 class="first-username">{{ item.username }}</h3>
          <!-- 一级评论发布时间 -->
          <span class="first-time">{{ item.date }}</span>
          <!-- 一级评论评论内容 -->
          <p class="first-comment">{{ item.content }}</p>
        </div>
      </div>
      <!-- 暂无评论的空状态 -->
      <el-empty description="暂无评论" v-if="show"></el-empty>
    </div>
	</div>

    <div class="head">
      <!-- 评论框 -->
      <input
        type="text"
        placeholder="请输入评论 . . ."
        ref="input"
        @focus="obtain"
        @blur="lose"
        v-model="firstComments"
        @keyup.enter="sumbit"
      />
      <!-- 发布按钮 -->
      <button @click="sumbit">发表评论</button>
    </div>
  </div>
</template>

<script>
export default {
  onLoad(options) {
    // 获取页面跳转参数
    this.video_id = options.video_id;
	this.avatar = uni.getStorageSync('avatar');
    this.token = uni.getStorageSync('token');
    console.log('评论界面接收到的video_id：', this.video_id);
    console.log('评论界面接收到的token：', this.token);
	 console.log('评论界面接收到的avatar：', this.avatar);

    // 在页面加载时调用获取评论列表的方法
    this.fetchComments();
  },
  data() {
    return {
      current: new Date(), // 时间
	  avatar:'',
      token: '',
      video_id: 0,
      picture: '',
      firstComments: '',
      comments: [],
      show: false, // 暂无条件显示
      item: {
        avatar: '',
      },
    };
  },
  methods: {
    obtain() {},
    // 评价框失焦
    lose() {},

// 获取评论列表的方法
fetchComments() {
  // 使用uni.request或者其他适合小程序的请求方法向后端发起获取评论列表的请求
  // 示例：
  uni.request({
    url: `http://192.168.47.9:8080/douyin/comment/list/?video_id=${this.video_id}&token=${this.token}`,
    method: 'GET',
    success: (res) => {
      // 检查comment_list是否存在
      if (res.data && res.data.comment_list) {
        // 请求成功后将评论列表赋值给comments数组
        this.comments = res.data.comment_list.map(comment => {
          return {
            username: comment.user.name,
            date: comment.create_date,
            content: comment.content,
            avatar: comment.user.avatar, // 头像字段，请根据实际返回的字段调整
            // 其他字段根据后端返回的数据结构进行调整
          };
        });
      } else {
        console.error('获取评论列表失败，comment_list不存在', res.data);
      }
    },
    fail: (err) => {
      console.error('获取评论列表失败', err);
    },
  });
},

      // 发表评论
      sumbit() {
        if (this.firstComments.trim() === '') {
          // 如果评论内容为空，不进行提交
          return;
        }
  
        // 发起请求到后端，提交评论
        uni.request({
          url: `http://192.168.47.9:8080/douyin/comment/action/?token=${this.token}&video_id=${this.video_id}&action_type=1&comment_text=${this.firstComments}`,
          method: 'POST',
          success: (res) => {
            if (res.data.status_code === 1) {
				// 如果后端返回成功，则刷新评论列表
				this.fetchComments();
              // 如果后端返回成功，则将新评论添加到comments数组
              const newComment = {
                username: 'Champion',
                date: this.formatDate(new Date()), // 使用一个辅助函数格式化日期
                content: this.firstComments,
                avatar: this.avatar, // 使用当前用户的头像
              };
              this.comments.unshift(newComment); // 将新评论插入到数组的开头
              this.firstComments = ''; // 清空输入框内容
            } else {
              console.error('评论发布失败', res.data);
            }
          },
          fail: (err) => {
            console.error('评论发布失败', err);
          },
        });
      },
  
      // 辅助函数，格式化日期
      formatDate(date) {
        return `${date.getFullYear()}年${date.getMonth() + 1}月${date.getDate()}日 ${date.getHours()}:${date.getMinutes()}:${date.getSeconds()}`;
      },
    },
  created() {},
};
</script>



<style>
	  /* 评论列表容器，固定高度，使得评论列表可以在容器内滚动 */
	  .comments-container {
	    max-height: 600px; /* 设置固定高度，可以根据实际情况调整 */
	    overflow-y: auto; /* 允许垂直滚动 */
	  }
	
	  /* 评论框 */
	  .head {
	    position: fixed; /* 固定在屏幕底部 */
	    bottom: 0; /* 距离底部的距离为0 */
	    background-color: rgb(248, 248, 248);
	    width: 100%; /* 宽度占满屏幕 */
	    height: 75px;
	    border-radius: 5px;
	  }
	/* 评论框 */
	.head input {
	  position: absolute;
	  top: 13px;
	  left: 20px;
	  height: 45px;
	  border-radius: 5px;
	  outline: none;
	  width: 60%;
	  font-size: 20px;
	  padding: 0 20px;
	  border: 2px solid #f8f8f8;
	   text-align: left; /* 文本左对齐 */
	}
	/* 发布评论按钮 */
	.head button {
	  position: absolute;
	  top: 13px;
	  right: 20px;
	  width: 120px;
	  height: 48px;
	  border: 0;
	  border-radius: 5px;
	  font-size: 20px;
	  font-weight: 500;
	  color: #fff;
	  background-color: rgb(118, 211, 248);
	  cursor: pointer;
	  letter-spacing: 2px;
	}
	/* 鼠标经过字体加粗 */
	.head button:hover {
	  font-weight: 600;
	}
	 
	/* 评论内容区域 */
	.content .first {
	  display: flex;
	  position: relative;
	  padding: 10px 10px 0px 0;
	  text-align: left;
	}
	.first .first-img {
	  flex: 1;
	  text-align: center;
	}
	.first img {
	  width: 50px;
	  height: 50px;
	  border-radius: 50%;
	}
	.first-username,
	.second-username {
	  color: #504f4f;
	}
	.first-content {
	  flex: 9;
	}
.first-time,
.second-time {
  color: #767575;
  text-align: right; /* 设置文本右对齐 */
  float: right; /* 设置元素向右浮动 */
   margin-top: 30px;
}

	.first-comment,
	.second-comment {
	  margin-top: 20px;
	}
	/* 删除评论 */
	.delete:hover {
	  color: red;
	}
	/* 评论字体图标 */
	.comments::before {
	  /* 想使用的icon的十六制编码，去掉&#x之后的 */
	  content: "\e8b9";
	  /* 必须加 */
	  font-family: "iconfont";
	  margin-right: 4px;
	  font-size: 20px;
	}

</style>
