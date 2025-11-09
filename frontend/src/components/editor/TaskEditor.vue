<script setup lang="ts">
import {useSessionStore} from '@/stores/session';

const session = useSessionStore()
export interface Form {
  name: string;
  profile: string;
  time: [Date, Date];
  statusId: number;
}

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
    <el-form-item :label="$t('time')">
      <el-date-picker v-model="model.time" type="datetimerange">
      </el-date-picker>
    </el-form-item>
    <el-form-item :label="$t('status')">
      <el-select v-model="model.statusId" :empty-values="[0]">
        <el-option
          v-for="s in session.enums['task_statuses']"
          :key="s.id"
          :label="s.name"
          :value="s.id"
        />
      </el-select>
    </el-form-item>
  </el-form>
</template>
