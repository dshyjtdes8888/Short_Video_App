<template>
  <div>
    <button
      @click="favor"
      class="my_button"
      :style="{ backgroundColor: bg_color, color: ft_color }"
      @mouseover="change"
      @mouseout="goback"
    >
      {{ content }}
    </button>
  </div>
</template>

<script>
import FollowListVue from '../mark/FollowList.vue';
export default {
  name: "zan",
    props: {
      user: {
        type: Object,
        required: true,
      },
      index: {
        type: Number,
        required: true,
      },
	  token:{
		type: String,
		required: true,
	  }
    },
  data() {
    return {
      liked: true, // Initial state is "取消关注"
      content: "x取消关注",
      bg_color: "#0697d9", // Initial color is red
      ft_color: "#ffffff",
    };
  },
  methods: {
    async favor() {
     // this.liked = !this.liked;
      if (!this.liked) {
        this.content = "x取消关注";
        this.bg_color = "#0697d9";
        this.ft_color = "#ffffff";
		    // 向后端发送关注状态变更请求
		    const response = await uni.request({
		      url: `http://192.168.47.9:8080/douyin/relation/action/?to_user_id=${this.user.id}&token=${this.token}&action_type=1`,
		      method: 'POST',
		      dataType: 'json',
		    });
			console.log("关注.",this.token);
		    if (response && response[1].statusCode === 200 && response[1].data.status_code === 0) {
				this.liked = !this.liked;
				console.log("关注/",this.token);
		    } else {
		      console.error('关注状态更新失败:', response[1].data ? response[1].data.status_msg : 'Unknown Error');
		    }
		  } else {
        this.content = "+关注";
        this.bg_color = "#f56c6c";
        this.ft_color = "#ffffff";
		// 向后端发送关注状态变更请求
		    const response = await uni.request({
		      url: `http://192.168.47.9:8080/douyin/relation/action/?to_user_id=${this.user.id}&token=${this.token}&action_type=2`,
		      method: 'POST',
		      dataType: 'json',
		    });
			
			console.log("取消关注",this.token);
			
		    if (response && response[1].statusCode === 200 && response[1].data.status_code === 0) {
				this.liked = !this.liked;
				console.log("取消关注",this.token);
		    } else {
		      console.error('关注状态更新失败:', response[1].data ? response[1].data.status_msg : 'Unknown Error');
		    }
		  } 
      },
	  
    change() {
      this.bg_color = "#0697d9";
      this.ft_color = "#ffffff";
    },
    goback() {
      if (this.liked) {
        this.bg_color = "#0697d9";
        this.ft_color = "#ffffff";
      } else {
        this.bg_color = "#f56c6c";
        this.ft_color = "#ffffff";
      }
    },
  },
};
</script>

<style>
button {
  outline: none;
}

.my_button {
  color: #f56c6c;
  background: #fef0f0;
  border: #fbc4c4 solid;
  border-radius: 20px;
  padding: 12px 23px;
  text-align: center;
  font-size: 16px;
  -webkit-transform: scale(0.7);
}
</style>
