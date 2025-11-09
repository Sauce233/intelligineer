<script setup lang="ts">
import {request} from '@/utils/axios';
import type {Attendance} from '@/utils/tables';
import {formatDate} from '@/utils/utils';
import {ArrowLeft} from '@element-plus/icons-vue';
import {ref} from 'vue';
import {useRoute} from 'vue-router';

const route = useRoute()

const attendance = ref<Attendance | null>(null)
request<Attendance>('GET', `/attendances/${route.params.id}`).then(res => attendance.value = res)
</script>

<template>

  <div class="p-4 flex flex-col gap-4">
    <div class="card flex gap-4">
      <el-button :icon="ArrowLeft" @click="$router.back">
        {{$t('back')}}
      </el-button>
      <div class="text-xl">{{$t('attendanceDetail.title')}}: {{attendance?.name}}</div>
    </div>
    <div class="text-xl">{{$t('attendanceDetail.infoTitle')}}</div>
    <div v-if="attendance" class="grid grid-cols-[repeat(auto-fill,minmax(250px,1fr))] gap-4">
      <div v-for="v in [
          ['name', attendance.name, ''],
          ['type', attendance.type.name, ''],
          ['createdAt', formatDate(attendance.createdAt), ''],
          ['startTime', formatDate(attendance.startTime), ''],
          ['endTime', formatDate(attendance.endTime), ''],
          ['location', attendance.location, 'col-span-3'],
        ]" :class="['flex items-center gap-2', v[2]]">
        <div class="text-subtle">{{$t(v[0])}}</div>
        <div>{{v[1]}}</div>
      </div>
    </div>

    <router-view />

  </div>
</template>
