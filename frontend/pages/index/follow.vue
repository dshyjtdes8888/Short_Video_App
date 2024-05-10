
<template>
	<view class="my">
	
		<view class="banner" :style="{ 'background-image': 'url(' + backgroundImageUrl + ')' }">
    </view>
		
		<view class="user-info grid col-2 align-center justify-between">
			<view class="header align-center justify-center">
				 <image :src="avatar"></image>
			</view>
      <view class="grid col-3 align-center justify-center text-white" style="width: 72%;">
        <view class="flex-direction align-center justify-center text-center">
          <view class="text-gray">获赞</view>
          <view class="text-bold text-xl">{{ likesCount }}</view>
        </view>
        <view class="flex-direction align-center justify-center text-center" @tap="goToFollows">
          <view class="text-gray">关注</view>
          <view class="text-bold text-xl">{{ followingsCount }}</view>
        </view>
        <view class="flex-direction align-center justify-center text-center" @tap="goToFollows">
          <view class="text-gray">粉丝</view>
          <view class="text-bold text-xl">{{ followersCount }}</view>
        </view>
      </view>
      
		</view>
		
		<!-- username -->
		<view class="grid user-detail col-2 align-center">
			<view class="left flex justify-center">
				<view class="">
					<text class="text-white text-xl text-bold">{{ username }} </text>
				</view>
				<view class="number align-center">
					<text class="text-gray text-df"> 抖音号：629901001 </text>
          <image src="/static/code.png"></image>
				</view>
			</view>
		</view>
    
    <view class="introduce">
      <view class="qianming">
      	<text class="text-white">个性签名：{{signature}}</text>
      </view>
<!--      <view class="x_int text-gray">
      	<text class="btn text-sm">北京</text>
        <text class="btn text-sm"> + 添加联系方式</text>
      </view> -->
    </view>
    
<view class="btn">
  <view v-if="isFollowed" class="button gray-button" @click="handleFollow">已关注</view>
  <view v-else class="button" @click="handleFollow">+关注</view>
</view>



		
		<!-- tab -->
		<scroll-view scroll-x class="nav text-center">
			<view class="cu-item text-gray" :class="index==TabCur?'text-white cur':''" v-for="(item,index) in tabList" :key="index" @tap="tabSelect" :data-id="index">
				{{item.name}} {{ item.number }}
			</view>
		</scroll-view>
		
		<!-- list -->
		    <view v-if="TabCur === 0">
		      <video-list :dataList="videoList"></video-list>
		    </view>
		    <view v-else-if="TabCur === 1">
		      <video-list :dataList="favoriteList"></video-list>
		    </view>
		
	</view>
</template>

<script>
	import videoList from '@/components/my/list.vue';
	//import favoriteList from '@/components/my/list.vue';
	export default {
		components: {
			videoList
		},
		data() {
			return {
				// work_count: 0,
				// favorite_count: 0,
				cover_url:'',
				backgroundImageUrl: '', // 存储背景图片的 URL
				signature:'',
				avatar: '/static/logo.jpg', // 默认头像
				user_id:0,
				token:0,
				likesCount: 0,
				followingsCount: 0,
				followersCount: 0,
				username: '',
				TabCur: 0,
				scrollLeft: 0,
				
				isFollowed: false,  // 设置为 false
				tabList:[
					{
						name:'作品',
						number: 0
					},{
						name:'喜欢',
						number: 0
					},
				],
				videoList: [],
				favoriteList:[],
				action_type:1,
        statusBarHeight:0, 
		user:[],
			};
		},
onLoad(options) {
  this.statusBarHeight = uni.getSystemInfoSync().statusBarHeight + 'px';
  
  //从本地存储中获取用户信息
  //    this.user_id = uni.getStorageSync('user_id');
	 // console.log('Token in onLoad:' ,this.user_id); // 在页面加载时输出 token 的值
  //    this.token = uni.getStorageSync('token');
	 // console.log('Token in onLoad:' ,this.token); // 在页面加载时输出 token 的值
	 // this.avatar = uni.getStorageSync('avatar');
	 // console.log('Token in onLoad:' ,this.user_id); // 在页面加载时输出 token 的值

	//console.log("author结构体",this.user)
	var data = JSON.parse(options.author);
	this.user = data;
	console.log("author结构体", options.author)
},

methods: {
  tabSelect(e) {
    this.TabCur = e.currentTarget.dataset.id;
    this.scrollLeft = (e.currentTarget.dataset.id - 1) * 60;
  },
  
  
async handleFollow() {
  try {
    const action_type = this.isFollowed ? 2 : 1; // 如果已关注，则取消关注；如果未关注，则关注
    const response = await uni.request({
      url: `http://192.168.47.9:8080/douyin/relation/action/?to_user_id=${this.user.user_id}&token=${this.user.username}&action_type=${action_type}`,
      method: 'POST',
      dataType: 'json',
    });

    console.log('关注Response:', response);

    if (response[1].data.status_code === 0) {
      // 成功处理
      this.isFollowed  = !this.isFollowed ; // 切换关注状态
    } else {
      console.error('关注失败:', response[1].data.status_msg || 'Unknown Error');
    }
  } catch (error) {
    console.error('关注失败:', error);
  }
},

	  
  async fetchData() {
    try {
      const response = await uni.request({
        url: `http://192.168.47.9:8080/douyin/user/?user_id=${this.user.user_id}&token=${this.user.username}`, // 替换为你的实际API URL
        method: 'GET',
        dataType: 'json',
      });

      const data = response[1].data;
      this.likesCount = data.user.total_favorited;
      this.followingsCount = data.user.follow_count;
      this.followersCount = data.user.follower_count;
      this.username = data.user.name;
	  this.signature = data.user.signature;
	  this.avatar = data.user.avatar || '/static/logo.jpg'; // 如果头像为空，则使用默认头像
	  this.backgroundImageUrl = data.user.background_image;
	  this.tabList[0].number = data.user.work_count;
	  this.tabList[1].number = data.user.favorite_count;
	  
	  
	   uni.setStorageSync('avatar', data.user.avatar || '/static/logo.jpg');
	  console.log(this.data)
      // 使用获取到的用户名设置页面标题
      uni.setNavigationBarTitle({
        title: this.username
      });
	  
	  
	  const response2 = await uni.request({
	    url: `http://192.168.47.9:8080/douyin/publish/list/?user_id=${this.user.user_id}&token=${this.user.username}`, // 替换为你的实际API URL
	    method: 'GET',
	    dataType: 'json',
	  });
	  
	  console.log("Response2",response2)
	  
       if (response2[1].data.status_code === 0) {
          this.videoList = response2[1].data.video_list;
        } else {
          console.error('获取发布视频列表失败:', response2[1].data.status_msg || 'Unknown Error');
        }
		
		
		
		// 获取喜欢视频列表
		        const response3 = await uni.request({
		          url: `http://192.168.47.9:8080/douyin/favorite/list/?user_id=${this.user.user_id}&token=${this.user.username}`,
		          method: 'GET',
		          dataType: 'json',
		        });
		
		        if (response3[1].data.status_code === 0) {
		          this.favoriteList = response3[1].data.video_list;
		        } else {
		          console.error('获取喜欢视频列表失败:', response3[1].data.status_msg || 'Unknown Error');
		        }
				
				
				

					
					
    } catch (error) {
      console.error('获取用户信息失败:', error);
    }
  },
  // 跳转到关注列表页面
    goToFollows() {
		console.log('点击了关注或者粉丝'); // 检查是否输出
      uni.navigateTo({
       // url: `/pages/mark/f?user_id=${this.data.user_id}&token=${this.data.token}`,
	   url: `/pages/mark2/f`,
      });
    },
},

mounted() {
  this.fetchData(); // 在组件挂载后调用方法获取数据
},

	}
</script>

<style>
page{
  background-color: #161823;
}
</style>
<style lang="less" scoped>
  .cu-item{
    width: 20%;
    font-size: 30upx;
  }
.my{
	height: 100vh;
	width: 100vw;
}
.banner{
	// background-image: url(../../static/my_bg.jpg);
	width: 100vw;
	min-height: 280upx;
	background-size: cover;
  position: relative;
  .left{
    position: absolute;
    left: 20upx;
    width: 60upx;
    height: 60upx;
    display: flex;
    align-items: center;
    justify-content: center;
    border-radius: 50%;
    background-color: #161823;
    opacity: .8;
    image{
      width: 36upx;
      height: 36upx;
    }
  }
  .right{
    position: absolute;
    right: 20upx;
    width: 60upx;
    height: 60upx;
    display: flex;
    align-items: center;
    justify-content: center;
    border-radius: 50%;
    background-color: #161823;
    opacity: .7;
    image{
      width: 36upx;
      height: 36upx;
    }
  }
}
.user-info{
  width: 94%;
  margin: 0 3%;
  position: relative;
	.header{
    background: #161823;
    width: 174upx;
    height: 174upx;
    display: flex;
    align-items: center;
    justify-content: center;
    margin-top: -60upx;
    border-radius: 50%;
    
		image{
			width: 160upx;
			height: 160upx;
      border-radius: 50%;
		}
	}
}

.btn {
  width: 100%;
  display: flex;
  justify-content: center;
  
}

.button {
  background: #e74c3c;
  width: 100%;
  text-align: center;
  height: 85rpx;
  line-height: 76rpx;
  border-radius: 8rpx;
  border: 2rpx solid #c0392b; /* 边框颜色，可根据需要调整 */
  box-sizing: border-box; /* 防止边框撑大元素 */
  overflow: hidden; /* 防止边框溢出 */
  color:white;
}

.gray-button {
  background: #ccc;
  border-color: #999; /* 灰色背景的边框颜色，可根据需要调整 */
  cursor: not-allowed;
}



.user-detail{
	border-bottom: 1px solid rgba(110, 111, 117, 0.5);
	padding: 10upx 0 15upx;
  width: 94%;
  margin: 0 3%;
	.left{
		flex-direction: column;
		.number{
			margin: 10upx 0;
      display: flex;
      image{
        width: 36upx;
        height: 36upx;
        margin-left: 16upx;
      }
		}
	}
}

.introduce{
  width: 94%;
  margin: 0 auto;
  padding: 20upx 0;
  .x_int{
    margin: 16upx 0;
    .btn{
      background: #3A3A44;
      padding: 6upx 12upx;
      margin-right: 10upx;
      border-radius: 6upx;
    }
  }
}

.adv{
	padding: 20upx 0;
  width: 94%;
  
  margin: 0 3%;
  image{
    width: 100%;
    height: 100%;
  }
}

.nav .cu-item.cur{
  border-bottom: 6rpx solid;
  border-color: #fbc732;
}
</style>
