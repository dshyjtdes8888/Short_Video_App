<template>
  <div>
    <div class="tab-bar">
      <span :class="{ 'active-tab': activeTab === 'following' }" @click="switchTab('following')">关注</span>
      <span :class="{ 'active-tab': activeTab === 'followers' }" @click="switchTab('followers')">粉丝</span>
    </div>
    
    <div v-if="activeTab === 'following'">
      <FollowList :list="followingList" @toggleFollow="toggleFollow" />
    </div>

    <div v-if="activeTab === 'followers'">
      <FollowersList :list="followersList" @toggleFollow="toggleFollow" />
    </div>
  </div>
</template>

<script>
import FollowList from 'pages/mark/FollowList'; // 你的关注列表组件路径
import FollowersList from 'pages/mark/FollowerList'; // 你的粉丝列表组件路径

export default {
  data() {
    return {
		followersList:[],
		followingList:[],
      activeTab: 'following',
    };
  },
  components: {
    FollowList,
    FollowersList,
  },
  methods: {




    switchTab(tab) {
      this.activeTab = tab;
    },
    toggleFollow(index) {
      const list = this.activeTab === 'following' ? 'followingList' : 'followersList';
      this.$set(this[list], index, {
        ...this[list][index],
        isFollowing: !this[list][index].isFollowing,
      });
    },
  },
  
  // mounted() {
  //   this.fetchData();
  // },
};
</script>

<style scoped>
.tab-bar {
  display: flex;
  justify-content: space-around;
  margin-bottom: 10px;
}

.tab-bar span {
  font-size: 16px;
  cursor: pointer;
}

.active-tab {
  color: #3498db;
  border-bottom: 2px solid #3498db;
}
</style>
