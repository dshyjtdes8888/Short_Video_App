
<template>
	<view class="msg">
		<!-- tools -->
		<view class="cu-list grid col-4 tools align-center justify-center">
				<view class="flex align-center justify-center">
					<image src="../../static/msg/fans.png" mode="widthFix"></image>
					<text>粉丝</text>
				</view>
				<view class="flex align-center justify-center">
					<image src="../../static/msg/zan.png" mode="widthFix"></image>
					<text>赞</text>
				</view>
				<view class="flex align-center justify-center">
					<image src="../../static/msg/at.png" mode="widthFix"></image>
					<text>@我的</text>
				</view>
				<view class="flex align-center justify-center">
					<image src="../../static/msg/comment.png" mode="widthFix"></image>
					<text>评论</text>
				</view>
		</view>
		
<view class="cu-list menu-avatar msg-list">
  <view class="cu-item bottom_line" v-for="(item, index) in list" :key="index" @click ="navigateToPersonalPage(item)">
    <view class="cu-avatar round lg" :style="'background-image:url('+item.avatar+');'"></view>
    <view class="content">
      <view class="text-grey">{{item.username}}</view>
      <view class="text-gray text-sm flex">
        <view class="text-cut">
          <text class="cuIcon-infofill text-red margin-right-xs" v-if="!item.success"></text>
          {{item.content}}
        </view>
      </view>
    </view>
    <view class="action">
      <view class="text-grey text-xs">{{item.time}}</view>
      <view class="cu-tag round bg-grey sm">{{item.number}}</view>
    </view>
  </view>
</view>

		
	</view>
</template>

<script>
	export default {
		data() {
			return {
				//name:'',///
				user_id: uni.getStorageSync('user_id'), // 从本地存储中获取头像信息
				token: uni.getStorageSync('token'), // 从本地存储中获取token信息
				
			list: [],
			
			};
		},
		 mounted() {
		    // 调用后端接口获取数据
		    this.fetchData();
		  },
		  
		  methods: {
			      // 使用 uniapp 的 navigateTo API 进行页面跳转
			      navigateToPersonalPage(item) {
			        // 使用 uniapp 的 navigateTo API 进行页面跳转，并传递参数到目标页面
			        uni.navigateTo({
			          url: `/pages/msg/msg_personal?username=${item.username}&avatar=${item.avatar}&to_user_id=${item.to_user_id}`, // 传递 item 的相关属性作为参数
			        });
			      },

		      async fetchData() {
		        try {
		          // 调用后端接口获取数据
		          const response = await uni.request({
		            url: `http://192.168.47.9:8080/douyin/relation/friend/list/?user_id=${this.user_id}&token=${this.token}`, // 替换为你的实际API URL
		            method: 'GET',
		            dataType: 'json',
		          });
		  
		          const data = response[1].data;
		          if (data.status_code === 0) {
		            // 处理后端返回的数据
		            this.list = data.user_list.map(user => ({
		              avatar: user.avatar,
		              username: user.name,
		              content: '你原来的内容属性', // 保持其他属性不变
		              time: '22.50', // 保持其他属性不变
		              number: '5', // 保持其他属性不变
		              success: true ,// 这里你需要判断用户是否成功
					  to_user_id:user.id,
		            }));
		          } else {
		            console.error('获取数据失败:', data.status_msg || 'Unknown Error');
		          }
		        } catch (error) {
		          console.error('获取数据失败:', error);
		        }
		      }
		    }
	}
</script>

<style lang="less">
	page{
    background-color: #161823;
  }
.msg{
	width: 100vw;
	height: 100vh;
	background-color: #161823;
}
.tools{
	view{
		display: flex;
		flex-direction: column;
	}
	image{
		width: 100upx;
		height: 100upx;
		display: flex;
	}
}
.cu-list.grid{
	background-color: #161823;
	padding: 30upx 0;
	color: #FFFFFF;
	border-bottom: 1px solid rgba(110,111,117,0.2);
}

.cu-list.menu-avatar>.cu-item:after,
.cu-list.menu>.cu-item:after{
  border-bottom-color: rgba(110,111,117,0.2) !important;
  border-bottom-width: 1px;
}
.cu-list+.cu-list{
	margin-top: 0;
}
.cu-list.menu-avatar>.cu-item{
	background-color: #161823;
	border-color: rgba(110,111,117,0.5);
}
</style>
