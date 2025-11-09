<script setup lang="ts">

export interface Form {
  title: string;
  content: string;
  minutes: number;
}

const model = defineModel<Form>({ required: true })
</script>

<template>
  <el-form :model="model">
    <el-form-item :label="$t('title')">
      <el-input v-model="model.title" class="!w-64" />
    </el-form-item>
    <el-form-item :label="$t('content')">
      <el-input v-model="model.content" type="textarea" />
    </el-form-item>
    <el-form-item :label="$t('minutes')">
      <el-input-number v-model="model.minutes" />
    </el-form-item>
  </el-form>
</template>
