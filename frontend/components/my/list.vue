<template>
	<view>
		<view class="list grid col-3 align-center justify-center">
			<block v-for="(item, index) in dataList" :key="index">
				<view class="video">
					<image :src="item.cover_url"></image>
          <view class="like align-center">
            <image src="/static/img/index/xin-2.png"></image>
            <text>{{item.favorite_count}}</text>
          </view>
				</view>
			</block>
		</view>
	</view>
</template>

<script>
export default {
	props: {
		dataList: {}
	},
	data() {
		return {
		};
	},
	// onLoad(){
	// 	this.cover_url = uni.getStorageSync('cover_url');
	// }
};
</script>

<style>
.list {
	display: flex;
	justify-content:left; 
	flex-wrap: wrap;
	padding-bottom: 100rpx;
/* 	background: #000000; */
	box-sizing: border-box;
  width: 100%;
}
.video{
  height: 380rpx;
  position: relative;
}
.grid.col-3>view {
	width: 250rpx;
  margin-bottom: 5rpx;
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
