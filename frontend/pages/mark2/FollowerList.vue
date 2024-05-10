<!-- FollowersList.vue -->
<!-- FollowersList.vue -->
<template>
  <div>
    <div class="followers-list-header">
      <h2>粉丝列表</h2>
    </div>
    <div class="followers-list">
      <div v-for="(user, index) in followersList" :key="index" class="follower-item">
        <img :src="user.avatar" alt="avatar" class="user-avatar">
        <div class="user-info">
          <h3>{{ user.name }}</h3>
          <p>{{ user.bio }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      followersList: [],
      user_id: 0,
      token: '',
    };
  },
  methods: {
    async fetchData() {
      try {
        // 获取粉丝列表数据
        const followersResponse = await uni.request({
          url: 'http://192.168.47.9:8080/douyin/relation/follower/list/',
          method: 'GET',
          dataType: 'json',
          data: {
            user_id: this.user_id,
            token: this.token,
          },
        });

        console.log('Followers Response:', followersResponse); // 输出响应信息

        if (followersResponse && followersResponse[1].statusCode === 200 && followersResponse[1].data.status_code === 0) {
          console.log('获取粉丝列表成功:', followersResponse[1].data.user_list);
          this.followersList = followersResponse[1].data.user_list.map(user => ({
            name: user.name,
            avatar: user.avatar,
            bio: user.signature,
            // 由于是粉丝列表，没有关注状态字段
          }));
        } else {
          console.error('获取粉丝列表失败:', followersResponse[1].data ? followersResponse[1].data.status_msg : 'Unknown Error');
        }

      } catch (error) {
        console.error('获取数据失败:', error);
      }
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
.followers-list-header {
  padding: 20px;
  text-align: center;
  background-color: #3498db;
  color: #fff;
}

.followers-list {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-around;
  padding: 20px;
}

.follower-item {
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
</style>
