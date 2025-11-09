<script setup lang="ts">

export interface Form {
  name: string;
  password: string;
  email: string;
  phone: string;
  profile: string;
}

const model = defineModel<Form>({ required: true })
</script>

<template>
  <el-form :model="model">
    <el-form-item :label="$t('name')">
      <el-input v-model="model.name" class="!w-64" />
    </el-form-item>
    <el-form-item :label="$t('password')">
      <el-input v-model="model.password" class="!w-64" type="password" />
    </el-form-item>
    <el-form-item :label="$t('email')">
      <el-input v-model="model.email" class="!w-64" />
    </el-form-item>
    <el-form-item :label="$t('phone')">
      <el-input v-model="model.phone" class="!w-64" />
    </el-form-item>
    <el-form-item :label="$t('profile')">
      <el-input v-model="model.profile" type="textarea" />
    </el-form-item>
  </el-form>
</template>
