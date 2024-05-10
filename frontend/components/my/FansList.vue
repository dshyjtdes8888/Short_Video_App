<template>
  <view>
    <!-- 用户列表 -->
    <view v-for="(user, index) in userList" :key="index" class="user-item">
      <!-- 用户头像 -->
      <image :src="user.avatar" class="avatar"></image>

      <!-- 用户信息 -->
      <view class="user-info">
        <text class="user-name">{{ user.username }}</text>
        <text class="user-description">{{ user.description }}</text>
      </view>
    </view>
  </view>
</template>

<script>
export default {
  props: {
    userList: {
      type: Array,
      default: () => [],
    },
  },
};
</script>

<style scoped>
/* 样式可以根据实际需要进行调整 */
.user-item {
  display: flex;
  align-items: center;
  padding: 10px;
  border-bottom: 1px solid #eee;
}

.avatar {
  width: 50px;
  height: 50px;
  border-radius: 25px;
  margin-right: 10px;
}

.user-info {
  flex: 1;
}

.user-name {
  font-size: 16px;
  font-weight: bold;
}

.user-description {
  font-size: 14px;
  color: #888;
}
</style>
