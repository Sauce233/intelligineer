<script setup lang="ts">

export interface Form {
  name: string;
  address: string;
  profile: string;
}

const model = defineModel<Form>({ required: true })
</script>

<template>
  <el-form v-model="model" label-width="auto">
    <el-form-item :label="$t('name')">
      <el-input v-model="model.name">
      </el-input>
    </el-form-item>
    <el-form-item :label="$t('address')">
      <el-input v-model="model.address" type="textarea">
      </el-input>
    </el-form-item>
    <el-form-item :label="$t('profile')">
      <el-input v-model="model.profile" type="textarea">
      </el-input>
    </el-form-item>
  </el-form>
</template>
