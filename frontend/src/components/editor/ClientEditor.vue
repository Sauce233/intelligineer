<script setup lang="ts">

export interface Form {
  name: string;
  address: string;
  profile: string;
}

const model = defineModel<Form>({ required: true })
</script>

<template>
  <el-form :model="model">
    <el-form-item :label="$t('name')">
      <el-input v-model="model.name" class="!w-64" />
    </el-form-item>
    <el-form-item :label="$t('address')">
      <el-input v-model="model.address" type="textarea" />
    </el-form-item>
    <el-form-item :label="$t('profile')">
      <el-input v-model="model.profile" type="textarea" />
    </el-form-item>
  </el-form>
</template>
