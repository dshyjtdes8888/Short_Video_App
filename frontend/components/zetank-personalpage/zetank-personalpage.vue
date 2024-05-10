<!-- 滑动切换选项卡+吸顶演示（待完善） -->
<template>
	<view class="content">
		<z-paging ref="paging" @scroll="scroll" :scrollable="scrollable" :hide-empty-view="true"
			:refresher-status.sync="refresherStatus" @query="queryList" :refresher-enabled="true">
			<view class="top-view">
				<image class="info-cover" @tap="fnMainBgPic" src="https://vkceyugu.cdn.bspapp.com/VKCEYUGU-5d36bb3d-8e98-4f00-ab8b-b018dc4f5dc6/17b0292c-8a63-4b7b-a6a6-724f416a9958.jpg"
				  mode="scaleToFill"></image>
				 <!-- 用户头像关注 -->
				 <view class="u-flex-wrap" style="padding-top: 18rpx;padding-bottom: 18rpx;position: relative;align-items: center;display: flex;flex-direction: row;justify-content: flex-end;">
				   <view class="info-avatar">
				     <u-avatar :src="userinfo.avatar" size="135"></u-avatar>
					 
				   </view>
					<view style="margin-right: 50rpx;">
						<u-tag v-if='showedit' text="编辑信息" mode="dark" shape="" bg-color="#f5991e" @click="UserEdit" />
						<view v-if='!showedit' style="height: 45rpx;"></view>
					</view>
				 </view>
				 <!-- 用户名 -->
				  <view style="display: flex;align-items: center;margin-left: 50rpx;margin-top: 20rpx;" v-if="userinfo.length!=0">
				    <text style="font-size: 32rpx;color: #111111;font-weight: bold;margin-right: 10rpx;">{{userinfo.nickname}}</text>
					<u-icon :name="userinfo.gender===1?'man':'woman'" :color="userinfo.gender===1?'#4baed5':'#f5991e'" size="30"></u-icon>
				  </view>
				  
				  <!-- 单位职务 -->
				  <view class="u-flex">
					  <view style="display: flex;align-items: center;margin-left: 50rpx;margin-top: 20rpx;" v-if="userinfo.length!=0">
					      <u-icon name="map-fill" size="32" color="#4baed5"></u-icon>
					      <text style="font-size: 24rpx;color: #555555;margin-left: 10rpx;">{{userinfo.city}}</text>
					    </view>
					  <view style="display: flex;align-items: center;margin-left: 50rpx;margin-top: 20rpx;" v-if="userinfo.length!=0">
					    <u-icon name="man-add-fill" size="32" color="#4baed5"></u-icon>
					    <text style="font-size: 24rpx;color: #555555;margin-left: 10rpx;">{{userinfo.jor}}</text>
					  </view>
				  </view>
				  <!-- 简介-->
				<view style="font-size: 24rpx;color: #555555;margin:20rpx 50rpx;" v-if="userinfo.length!=0">
					<u-tag v-if='membertype!=""' :text="membertype" mode="dark" shape="circle" bg-color="#4baed5" />
					<text style="margin-right: 20rpx;">{{Information==''?'暂无简介':Information}}</text>
				</view>
				<view style="display: flex;align-items: center;margin-left: 50rpx;margin-top: 20rpx;margin-bottom: 20rpx;" v-if="userinfo.length!=0">
				  <text style="font-size: 32rpx;font-weight: bold;color: #555555;margin-right: 18rpx;">{{userinfo.good.length || 0}}</text><text style="font-size: 24rpx;color: gray;margin-right: 28rpx;">粉丝</text>
					<text style="font-size: 32rpx;font-weight: bold;color: #555555;margin-right: 18rpx;">{{userinfo.good.length || 0}}</text><text style="font-size: 24rpx;color: gray;margin-right: 28rpx;">关注</text>
					<text style="font-size: 32rpx;font-weight: bold;color: #555555;margin-right: 18rpx;">{{userinfo.good.length || 0}}</text><text style="font-size: 24rpx;color: gray;margin-right: 28rpx;">点赞</text>
				</view>
			</view>
			<view id="tabbar" class="paging-content" :style="'height:' + pageHeight + 'px'">
				<!-- 小程序中直接修改组件style为position: sticky;无效，需要在组件外层套一层view -->
				<view style="z-index: 100;position: sticky;top :0px;">
					<u-tabs-swiper ref="uTabs" :list="list" :current="current" @change="tabsChange" :is-scroll="false"
						swiperWidth="750"></u-tabs-swiper>
				</view>
				<swiper style="height: 100%;" :current="swiperCurrent" @transition="transition" @animationfinish="animationfinish">
					<swiper-item>
						<swiper-list-item ref="swiperItem" :refresher="false" :scrollable="itemscrollable" :tabIndex="0" :currentIndex="swiperCurrent" :fixed="false" :auto="true" @setScrollable="setScrollable"></swiper-list-item>
					</swiper-item>
					<swiper-item>
						
					</swiper-item>
				</swiper>
			</view>
		</z-paging>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				userinfo:{
					nickname:'zetank',
					avatar:"https://vkceyugu.cdn.bspapp.com/VKCEYUGU-b9e817e6-7ee6-461f-8e66-4bed243cbe98/ad090f5f-2a0f-4ea8-a8b5-fc8dafa4039e.jpg",
					gender:1,
					city:'广州',
					jor:'职务',
					good:{}
				},
				Information:'本人性格热情开朗,待人友好,为人诚实谦虚。工作勤奋,认真负责,能吃苦耐劳,尽职 尽责,有耐心。具有亲和力,平易近人,善于与人沟通',
				membertype:'',
				showedit:true,//信息编辑按钮
				refresherStatus: 0,
				scrollTopMap: {},
				// 页面高度
				pageHeight: 0,
				// header高度
				headerHeight: 0,
				scrollable: true,
				itemscrollable:false,
				change:false,
				list: [{
					name: '测试1'
				}, {
					name: '测试2'
				}, {
					name: '测试3'
				}, {
					name: '测试4'
				}],
				// 因为内部的滑动机制限制，请将tabs组件和swiper组件的current用不同变量赋值
				current: 0, // tabs组件的current值，表示当前活动的tab选项
				swiperCurrent: 0, // swiper组件的current值，表示当前那个swiper-item是活动的
			}
		},
		props:{
			uid:{
				type:String,
				default:""
			}
		},
		created() {
			this.pageHeight = uni.getSystemInfoSync().windowHeight;
			this.$nextTick(() => {
				const query = uni.createSelectorQuery().in(this);
				query.select('.top-view').boundingClientRect(data => {
					this.headerHeight = data.height;
				}).exec();
			})
		},
		methods: {
			queryList() {
				//触发了下拉刷新，通过当前tabIndex对应的列表下拉刷新
				this.$refs.swiperItem.reload();
				this.$refs.paging.complete([])
			},
			scroll(e) {
				let scrollTop = e.detail.scrollTop;
				//如果当前页面的scroll-view的scrollTop大于等于headerView的高度，则代表吸顶了
				if (scrollTop < this.headerHeight-0.5) {
					//还没吸顶
					//禁止子组件的z-paging(scroll-view)滚动，当前页面的z-paging(scroll-view)允许滚动
					this.scrollable = true;
					// this.$refs.swiperItem[this.current].setScrollable(false);
					this.itemscrollable=false
				} else {
					//吸顶了
					//允许子组件的z-paging(scroll-view)滚动，当前页面的z-paging(scroll-view)禁止滚动
					if(scrollTop >= this.headerHeight){
						this.scrollable = false;
					}
					// this.$refs.swiperItem[this.current].setScrollable(true);
					this.itemscrollable=true
				}
			},
			setScrollable(scrollable) {
				this.scrollable = scrollable;
			},
			// tabs通知swiper切换
			tabsChange(index) {
				this.swiperCurrent = index;
				this.$refs.paging.scrollToY(this.headerHeight)
			},
			// swiper-item左右移动，通知tabs的滑块跟随移动
			transition(e) {
				let dx = e.detail.dx;
				this.$refs.uTabs.setDx(dx);
			},
			// 由于swiper的内部机制问题，快速切换swiper不会触发dx的连续变化，需要在结束时重置状态
			// swiper滑动结束，分别设置tabs和swiper的状态
			animationfinish(e) {
				let current = e.detail.current;
				this.$refs.uTabs.setFinishCurrent(current);
				this.swiperCurrent = current;
				this.current = current;
			},
			// 修改用户背景封面图
			fnMainBgPic() {
				
			},
			//编辑信息
			UserEdit(){
				uni.navigateTo({
					url:""
				})
			},
		}
	}
</script>

<style>
	.banner-view {
		background-color: #007AFF;
		color: white;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
	}

	.paging-content {
		flex: 1;
		display: flex;
		flex-direction: column;
	}

	.swiper {
		flex: 1;
	}
	.info-cover {
		  display: block;
		  width: 100%;
		  height: 260rpx;
		}
	.info-avatar {
	  position: absolute;
	  left: 0;
	  top: -56rpx;
	  margin-left: 50rpx;
	}

</style>
