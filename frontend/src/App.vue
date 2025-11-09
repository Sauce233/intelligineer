<script setup lang="ts">
</script>

<template>
  <div class="bg-gray-50 dark:bg-gray-900 text-gray-900 dark:text-gray-100">
    <el-config-provider :button="{ round: true, plain: true, text: true }">
      <router-view>
      </router-view>
    </el-config-provider>
  </div>
</template>
