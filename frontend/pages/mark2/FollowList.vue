<!-- FollowList.vue -->
<template>
  <div>
    <div class="follow-list-header">
      <h2>关注列表</h2>
    </div>
    <div class="follow-list">
      <div v-for="(user, index) in followingList" :key="index" class="follow-item">
        <img :src="user.avatar" alt="avatar" class="user-avatar">
        <div class="user-info">
          <h3>{{ user.name }}</h3>
          <p>{{ user.bio }}</p>
        </div>
        <FollowButton :isFollowing="user.isFollowing" @toggleFollow="toggleFollow(index)" />
      </div>
    </div>
  </div>
</template>

<script>
import FollowButton from "pages/mark/FollowButton"; // 你的关注按钮组件路径

export default {
  data() {
    return {
      followingList: [],
      user_id: 0,
      token: '',
    };
  },
  components: {
    FollowButton,
  },
  methods: {
    async fetchData() {
      try {
        // 获取关注列表数据
        const followingResponse = await uni.request({
          url: 'http://192.168.47.9:8080/douyin/relation/follow/list/',
          method: 'GET',
          dataType: 'json',
          data: {
            user_id: this.user_id,
            token: this.token,
          },
        });

        console.log('Following Response:', followingResponse); // 输出响应信息

        if (followingResponse && followingResponse[1].statusCode === 200 && followingResponse[1].data.status_code === 0) {
          console.log('获取关注列表成功:', followingResponse[1].data.user_list);
          this.followingList = followingResponse[1].data.user_list.map(user => ({
            name: user.name,
            avatar: user.avatar,
            bio: user.signature,
            isFollowing: user.is_follow,
          }));
        } else {
          console.error('获取关注列表失败:', followingResponse[1].data ? followingResponse[1].data.status_msg : 'Unknown Error');
        }

      } catch (error) {
        console.error('获取数据失败:', error);
      }
    },

    toggleFollow(index) {
      this.$set(this.followingList, index, {
        ...this.followingList[index],
        isFollowing: !this.followingList[index].isFollowing,
      });
    },
  },

  created() {
    // 从本地存储中获取用户信息
    this.user_id = uni.getStorageSync('user_id');
    console.log('User_id in created:', this.user_id); // 在页面加载时输出 token 的值
    this.token = uni.getStorageSync('token');
    console.log('Token in created:', this.token); // 在页面加载时输出 token 的值
    this.fetchData(); // 调用 fetchData 方法获取数据
  },
};
</script>

<style scoped>
.follow-list-header {
  padding: 20px;
  text-align: center;
  background-color: #3498db;
  color: #fff;
}

.follow-list {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-around;
  padding: 20px;
}

.follow-item {
  width: 100%;
  margin-bottom: 20px;
  text-align: center;
  border: 1px solid #ddd;
  border-radius: 10px;
  padding: 10px;
  background-color: #fff;
}

.user-avatar {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  margin-bottom: 10px;
}

.user-info {
  margin-bottom: 10px;
}

.follow-button {
  padding: 5px 10px;
  font-size: 12px;
  cursor: pointer;
  border: none;
  border-radius: 5px;
}

.following {
  background-color: #e74c3c;
  color: #fff;
}

.following:hover {
  background-color: #c0392b;
}

.follow {
  background-color: #3498db;
  color: #fff;
}

.follow:hover {
  background-color: #2980b9;
}
</style>