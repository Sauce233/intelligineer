<script setup lang="ts">
import {useSessionStore} from '@/stores/session';

const session = useSessionStore()

export interface Form {
  name: string;
  location: string;
  typeId: string;
  free: boolean;
  time: [Date, Date],
}

const model = defineModel<Form>({ required: true })
</script>

<template>
  <el-form :model="model">
    <el-form-item :label="$t('name')">
      <el-input v-model="model.name" class="!w-64">
      </el-input>
    </el-form-item>
    <el-form-item :label="$t('location')">
      <el-input v-model="model.location" type="textarea" autosize>
      </el-input>
    </el-form-item>
    <el-form-item :label="$t('type')">
      <el-select v-model="model.typeId" :empty-values="[0]">
        <el-option
          v-for="v in session.enums['attendance_types']"
          :key="v.id"
          :label="v.name"
          :value="v.id"
        />
      </el-select>
    </el-form-item>
    <el-form-item :label="$t('free')">
      <el-switch v-model="model.free">
      </el-switch>
    </el-form-item>
    <el-form-item :label="$t('time')">
      <el-date-picker v-model="model.time" type="datetimerange">
      </el-date-picker>
    </el-form-item>
  </el-form>
</template>
