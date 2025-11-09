<script setup lang="ts">
import {useSessionStore} from '@/stores/session';


export interface Form {
  name: string;
  number: string;
  profile: string;
  amount: number;
  time: [Date, Date],
  statusId: number;
}

const model = defineModel<Form>({ required: true })
const session = useSessionStore()
</script>

<template>
  <el-form :model="model">
    <el-form-item :label="$t('name')">
      <el-input v-model="model.name" class="!w-64" />
    </el-form-item>
    <el-form-item :label="$t('number')">
      <el-input v-model="model.number" class="!w-64" />
    </el-form-item>
    <el-form-item :label="$t('profile')">
      <el-input v-model="model.profile" type="textarea" />
    </el-form-item>
    <el-form-item :label="$t('amount')">
      <el-input-number v-model="model.amount">
        <template #prefix>¥</template>
      </el-input-number>
    </el-form-item>
    <el-form-item :label="$t('time')">
      <el-date-picker v-model="model.time" type="datetimerange">
      </el-date-picker>
    </el-form-item>
    <el-form-item :label="$t('status')">
      <el-select v-model="model.statusId" :empty-values="[0]">
        <el-option
          v-for="v in session.enums['invoice_statuses']"
          :key="v.id"
          :label="v.name"
          :value="v.id"
        />
      </el-select>
    </el-form-item>
  </el-form>
</template>
