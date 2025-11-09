<script setup lang="ts">

export interface Form {
  name: string;
  email: string;
  phone: string;
  profile: string;
}

const model = defineModel<Form>({ required: true })
</script>

<template>
  <el-form :model="model" label-width="auto">
    <el-form-item :label="$t('name')">
      <el-input v-model="model.name" class="!w-64">
      </el-input>
    </el-form-item>
    <el-form-item :label="$t('email')">
      <el-input v-model="model.email" class="!w-64">
      </el-input>
    </el-form-item>
    <el-form-item :label="$t('phone')">
      <el-input v-model="model.phone" class="!w-64">
      </el-input>
    </el-form-item>
    <el-form-item :label="$t('profile')">
      <el-input v-model="model.profile" type="textarea">
      </el-input>
    </el-form-item>
  </el-form>
</template>
