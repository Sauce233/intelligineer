<script setup lang="ts">
import {request} from '@/utils/axios';
import type {UserAttendance} from '@/utils/tables';
import {formatDate, type List} from '@/utils/utils';
import {useUrlSearchParams} from '@vueuse/core';
import {useRouteParams} from '@vueuse/router';
import {ref, watchEffect} from 'vue';

const params = useUrlSearchParams<any>('history')
const id = useRouteParams('id')
const userAttendance = ref<List<UserAttendance> | null>(null)
watchEffect(async () => userAttendance.value = await request('GET', `/attendances/${id.value}/users`, params))

</script>

<template>
  <div class="text-xl">{{$t('attendanceDetail.usersTitle')}}</div>
  <div v-if="userAttendance" class="grid grid-cols-[repeat(auto-fill,minmax(250px,1fr))] gap-4">
    <div v-for="u in userAttendance.data" :key="u.user.id" class="card !p-0 flex flex-col gap-2">
      <img v-if="u.time" :src="`https://picsum.photos/seed/${Math.random()*100}/400/300`" class="h-32" alt="">
      <div v-else class="h-32"></div>
      <div class="flex gap-2 items-center p-4">
        <tag :color="u.status.color">{{u.status.name}}</tag>
        <span>{{u.user.name}}</span>
        <span v-if="u.time" class="text-subtle ms-auto">{{formatDate(u.time)}}</span>
      </div>
    </div>
  </div>

  <el-pagination
    v-if="userAttendance"
    layout="prev, pager, next, total"
    :total="userAttendance.total"
    :current-page="Number(params.page) || 1"
    @update:current-page="v => params.page = v"
    :page-size="Number(params.size) || 12"
    @update:page-size="v => params.size = v"
  />

</template>
