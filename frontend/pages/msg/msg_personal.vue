<template>
  <view class="chat-container">
    <!-- 顶部显示对方的头像和昵称 -->
    <view class="header">
      <image class="avatar" :src="this.avatar" mode="aspectFill"></image>
      <text class="username">{{this.username }}</text>
    </view>

    <!-- 中间部分显示聊天记录 -->
    <scroll-view class="chat-content" scroll-y>
      <!-- 我在右边，对方在左边 -->
      <view v-for="message in messages" :key="message.id" class="message-item" :class="{'my-message':message.from_user_id === user_id }"> <!-- currentUserID替换为自己id-->

		<view class="message-content-container">
		      <text class="message-content">{{ message.content }}</text>
		      <!-- <text v-if="message.create_time" class="message-time">{{ message.create_time }}</text> -->
		    </view>
		
      </view>
    </scroll-view>

    <!-- 底部输入框和发送按钮 -->
  <view class="footer">
      <input v-model="inputMessage" class="input-box" placeholder="请输入消息."></input>
      <button @click="sendMessage" class="send-button">发送</button>
    </view>
  </view>
</template>

<script>
export default {
  data() {
    return {
		user_id: uni.getStorageSync('user_id'), // 从本地存储中获取头像信息
		token: uni.getStorageSync('token'), // 从本地存储中获取token信息
		list: [],
	    to_user_id: 0,
		pre_msg_time:0,
	    username: '',
	    avatar: '',
      messages: [], // 存放聊天记录的数组
      inputMessage: '',
	  action_type:1,
    };
  },
  mounted() {
    // 在页面加载时调用获取聊天记录的接口
    this.getChatRecords();
  },
  
onLoad(options) {
  // 获取页面跳转参数
  this.user_id = uni.getStorageSync('user_id'); // 设置 user_id
  this.username = options.username;
  this.avatar = options.avatar;
  this.to_user_id = options.to_user_id;
  // 在控制台输出检查是否正确获取参数
  console.log('接收到的用户名：', this.username);
  console.log('接收到的头像URL：', this.avatar);
  console.log('接收到的to_user_id：', this.to_user_id);
  console.log('设置的user_id：', this.user_id);
  
},

  
  methods: {
	  
	  async getChatRecords() {
	        try {
	          const response = await uni.request({
	           url: `http://192.168.47.9:8080/douyin/message/chat/?to_user_id=${this.to_user_id}&token=${this.token}&pre_msg_time=${this.pre_msg_time}`, // 替换为你的实际API URL
	            method: 'GET',
	            dataType: 'json',
	          });
	  
	          if (response[1] && response[1].data && response[1].data.status_code === 0) {
	            // 获取聊天记录成功
				
				
	            this.messages = response[1].data.message_list.map(message => ({
					
	              ...message,
	              create_time: this.formatTime(message.create_time) // 格式化时间
	            }));
				console.log("messagge",this.messages)
	          } else {
	            console.error('获取聊天记录失败:', response);
	          }
	        } catch (error) {
	          console.error('Error during getChatRecords:', error);
	        }
	      },
		  
	      async sendMessage() {
	        try {
	          const response = await uni.request({
	            url: `http://192.168.47.9:8080/douyin/message/action/?to_user_id=${this.to_user_id}&token=${this.token}&action_type=${this.action_type}&content=${this.inputMessage}`, // 替换为实际的接口地址
	            method: 'POST',
	            dataType: 'json',
	            // data: {
	            //   token: 'your_auth_token', // 替换为实际的用户鉴权token
	            //   to_user_id: 'friend_user_id', // 替换为实际的对方用户id
	            //   action_type: 1, // 1 表示发送消息
	            //   content: this.inputMessage // 获取输入框中的消息内容
	            // }
	          });
	  
	          if (response[1] && response[1].data && response[1].data.status_code === 0) {
	            // 发送消息成功，刷新聊天记录
	            this.getChatRecords();
	            // 清空输入框
	            this.inputMessage = '';
	          } else {
	            console.error('发送消息失败:', response);
	          }
	        } catch (error) {
	          console.error('Error during sendMessage:', error);
	        }
	      },
		  
    formatTime(timestamp) {
      // 根据实际需求进行时间格式化
      // 示例：这里简单地将时间戳转换为字符串
      const date = new Date(timestamp);
      return date.toLocaleString();
    },
	
  }
};
</script>




<style scoped>
/* 样式部分，可以根据需要进行调整 */
.chat-container {
  display: flex;
  flex-direction: column;
  height: 100vh;
  background-color: #ffffff;
}

.header {
  display: flex;
  align-items: center;
  padding: 10px;
  background-color: #e6e6e6;
}

.avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  margin-right: 10px;
}

.username {
  font-size: 16px;
  color: #333;
}

.chat-content {
  flex: 1;
  padding: 10px;
}

.message-item {
  max-width: 80%;
  padding: 8px;
  margin-bottom: 10px;
  border-radius: 8px;
}

.message-item {
  display: flex;
  flex-direction: column;
  align-items: flex-start; /* 默认左对齐 */
  background-color: #bdbdbd;
  color: #060606;
  width: 30%;
}

.my-message.message-item {
  align-items: flex-end; /* 我的消息右对齐 */
  background-color: #4caf50;
  width: 30%;
  margin-left: 70%; /* 负值表示向右平移 */
}

.message-content-container {
  display: inline-block;
}

.message-content {
  color: #fff;
  padding: 15px;
  border-radius: 10px;
  font-size: 16px;
}

.message-time {
  margin-top: 4px;
  color: #666;
}

.footer {
  display: flex;
  align-items: center;
  padding: 10px;
  background-color: #e6e6e6;
}

.input-box {
  flex: 1;
  height: 40px;
  border: 1px solid #ccc;
  border-radius: 20px;
  padding: 0 10px;
  background-color: #fff;
}

.send-button {
  margin-left: 10px;
  padding: 0px 0px;
  border-radius: 10px;
  background-color: #46a2df;
  color: #fff;
  cursor: pointer;
  font-size: 16px;
}
</style>