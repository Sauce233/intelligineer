<script setup lang="ts">
import Tag from '@/components/Tag.vue';
import {request} from '@/utils/axios';
import type {Record, User} from '@/utils/tables';
import {formatDate, type List} from '@/utils/utils';
import {Back, Delete, Edit} from '@element-plus/icons-vue';
import {computed, reactive, ref, watchEffect} from 'vue';
import { use } from 'echarts/core';
import {
  GridComponent,
} from "echarts/components";
import { BarChart } from 'echarts/charts';
import { CanvasRenderer } from 'echarts/renderers';
import type { EChartsOption } from 'echarts';
import VChart from 'vue-echarts';
import type {Form} from '@/components/editor/UserEditor.vue';
import UserEditor from '@/components/editor/UserEditor.vue';
import {useRouteParams} from '@vueuse/router';

const id = useRouteParams('id')
use([GridComponent, BarChart, CanvasRenderer])

const workHoursData = ref<{
  date: string,
  total: number,
}[]>([])

request<any>('GET', `/user-hours/${id.value}/7`).then(res => workHoursData.value = res)

const workHoursOption = computed<EChartsOption>(() => ({
  xAxis: {
    type: 'category',
    data: ['週一', '週二', '週三', '週四', '週五', '週六', '週日'],
  },
  yAxis: {
    type: 'value',
    name: '分鐘' as any,
  },
  series: [
    {
      data: workHoursData.value.map(v => v.total),
      type: 'bar'
    }
  ]
}))

const user = ref<User | null>(null)
const load = async () => user.value = await request<User>('GET', `/users/${id.value}`)
load()

const records = ref<List<Record> | null>(null)

watchEffect(async () => records.value = await request('GET', '/records', { userId: id.value }))

const isDialogOpen = ref(false)
const form = reactive<Form>({
  name: '', password: '', email: '', phone: '', profile: '',
})

</script>

<template>
  <div class="p-4 flex flex-col gap-4">
    <div v-if="user" class="card flex gap-2 flex-wrap items-center">
      <el-button :icon="Back" @click="$router.back">
        {{$t('back')}}
      </el-button>
      <div class="text-lg">{{user.name}}</div>
      <tag v-for="r in user.roles" :key="r.id" :color="r.color">{{r.name}}</tag>
      <div class="text-subtle">{{user.email}}</div>
      <el-button :icon="Edit" circle class="ms-auto" @click="isDialogOpen = true" />
      <el-button :icon="Delete" circle />
    </div>

    <div class="flex justify-between">
      <div class="text-xl">{{$t('userDetail.workHoursTitle')}}</div>
    </div>

    <v-chart class="!h-72" :option="workHoursOption">
    </v-chart>

    <div class="text-2xl">{{$t('userDetail.infoTitle')}}</div>
    <div v-if="user" class="grid grid-cols-[repeat(auto-fill,minmax(250px,1fr))] gap-4">
      <div v-for="v in [
          ['createdAt', formatDate(user.createdAt), ''],
          ['updatedAt', formatDate(user.updatedAt), ''],
          ['name', user.name, ''],
          ['email', user.email, ''],
        ]" :class="['flex items-center gap-2', v[2]]">
        <div class="text-subtle">{{$t(v[0])}}</div>
        <div>{{v[1]}}</div>
      </div>
    </div>
    <div v-if="user" class="bg-subtle rounded-2xl p-4">
      {{user.profile}}
    </div>

    <div class="p-2 flex gap-1">
      <router-link
        v-for="v in [
          [`/users/${id}`, $t('profile')],
          [`/users/${id}/records`, $t('record')],
          [`/users/${id}/applications`, $t('application')],
          [`/users/${id}/attendances`, $t('attendance')],
        ]"
        :key="v[0]"
        :to="v[0]"
        class="p-2 rounded-2xl"
        exact-active-class="font-bold bg-subtle"
      >
        {{v[1]}}
      </router-link>
    </div>

    <router-view>
    </router-view>

  </div>

  <el-dialog v-model="isDialogOpen" :title="$t('user.detail.dialog.title')">
    <user-editor v-model="form">
    </user-editor>
    <template #footer>
      <el-button
        type="primary"
        @click="request('PATCH', '/users/${id}', form).then(ok => ok ?? load())"
      >
        {{$t('confirm')}}
      </el-button>
    </template>
  </el-dialog>

</template>
