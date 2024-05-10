<template>
	<view>
		<view class="list grid col-3 align-center justify-center">
			<block v-for="(item, index) in video_list" :key="index">
				<view class="video">
					<image :src="cover_url"></image>
          <view class="like align-center">
            <image src="/static/img/index/xin-2.png"></image>
            <text>{{like}}</text>
          </view>
				</view>
			</block>
		</view>
	</view>
</template>

<script>
export default {
	data() {
		return {
			like:0,
			cover_url:'',
			video_list:[],
		};
	},
	onload(){
		     this.user_id = uni.getStorageSync('user_id');
			 console.log('Token in created:' ,this.user_id); // 在页面加载时输出 token 的值
		     this.token = uni.getStorageSync('token');
			 console.log('Token in onLoad:' ,this.token); // 在页面加载时输出 token 的值
			console.log()
	},
	methods: {
	  tabSelect(e) {
	    this.TabCur = e.currentTarget.dataset.id;
	    this.scrollLeft = (e.currentTarget.dataset.id - 1) * 60;
	  },
	
	  async fetchData() 
	  {
	    try {
		  const response2 = await uni.request({
		    url: `http://192.168.47.9:8080/douyin/publish/list/?user_id=${this.user_id}&token=${this.token}`, // 替换为你的实际API URL
		    method: 'GET',
		    dataType: 'json',
		  });
		  
		  if (response[1].data.status_code === 0) {
			   uni.setStorageSync('cover_url', cover_url);
		            // 处理后端返回的视频列表数据
		            this.video_list = data.video_list;
					this.cover_url = data.video_list.cover_url;
					this.like = data.video_list.favorite_count;
		          } else {
		            console.error('获取视频列表失败:', data.status_msg || 'Unknown Error');
		          }
				  
	
	    } catch (error) {
	      console.error('获取用户信息失败:', error);
	    }
		}
		}
};
</script>

<style>
.list {
	display: flex;
	justify-content:space-between;
	flex-wrap: wrap;
	padding-bottom: 100rpx;
	background: #ffffff;
	box-sizing: border-box;
  width: 100%;
}
.video{
  height: 360rpx;
  position: relative;
}
.grid.col-3>view {
	width: 244rpx;
  margin-bottom: 10rpx;
}
.video image{
  width: 100%;
  height: 100%;
}
.video .like{
  position: absolute;
  left: 10rpx;
  bottom: 10rpx;
  color: #FFF;
  display: flex;
}
.video .like image{
  width: 40rpx;
  height: 40rpx;
  margin-right: 6rpx;
}
</style>
