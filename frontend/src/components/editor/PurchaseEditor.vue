<script setup lang="ts">
import {useSessionStore} from '@/stores/session';


export interface Form {
  name: string;
  profile: string;
  statusId: number;
}

const session = useSessionStore()
const model = defineModel<Form>({ required: true })

</script>

<template>
  <el-form :model="model">
    <el-form-item :label="$t('name')">
      <el-input v-model="model.name" class="!w-64" />
    </el-form-item>
    <el-form-item :label="$t('profile')">
      <el-input v-model="model.profile" type="textarea" />
    </el-form-item>
    <el-form-item :label="$t('status')">
      <el-select v-model="model.statusId" :empty-values="[0]">
        <el-option
          v-for="v in session.enums['purchase_statuses']"
          :key="v.id"
          :label="v.name"
          :value="v.id"
        />
      </el-select>
    </el-form-item>
  </el-form>
</template>
