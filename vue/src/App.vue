<template>
  <div class="box">
    <!-- 输入框只用来临时收集用户输入，不直接绑定到显示的 ref 上 -->
    <input type="text" v-model="inputVal1" placeholder="请输入第一个文本框" />
    <input type="text" v-model="inputVal2" placeholder="请输入第二个文本框" />
    
    <!-- 增加一个确定按钮，点击触发 handleClick 函数 -->
    <button @click="handleClick">确定</button>

    <!-- 这里才是真正用到 ref 变量的地方，只有点击确定后才会更新 -->
    <p>文本框1的内容是：{{ text1 }}</p>
    <p>文本框2的内容是：{{ text2 }}</p>
  </div>
</template>

<script setup>
import { ref } from 'vue'

// 这两个是最终用来显示的 ref 交互式变量（初始为空）
const text1 = ref('')
const text2 = ref('')

// 这两个是普通的 ref，仅仅作为输入框的临时“中转站”
const inputVal1 = ref('')
const inputVal2 = ref('')

// 点击按钮执行的函数
const handleClick = () => {
  // 点击确定时，才把中转站的值赋给真正的交互式变量
  text1.value = inputVal1.value
  text2.value = inputVal2.value
}
</script>

<style scoped>
.box {
  display: flex;
  flex-direction: column;
  gap: 10px;
}
input {
  padding: 8px;
  width: 200px;
  border: 1px solid #ccc;
  border-radius: 4px;
}
/* 给按钮加一点简单的样式，让好看一点 */
button {
  width: 208px; /* 和输入框对齐 */
  padding: 8px;
  background-color: #409eff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}
button:hover {
  background-color: #66b1ff;
}
</style>
