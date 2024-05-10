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
export default {
  name: "zan",
  data() {
    return {
      liked: true, // Initial state is "取消关注"
      content: "x取消关注",
      bg_color: "#0697d9", // Initial color is red
      ft_color: "#ffffff",
    };
  },
  methods: {
    favor() {
      this.liked = !this.liked;
      if (this.liked) {
        this.content = "x取消关注";
        this.bg_color = "#0697d9";
        this.ft_color = "#ffffff";
      } else {
        this.content = "+关注";
        this.bg_color = "#f56c6c";
        this.ft_color = "#ffffff";
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
