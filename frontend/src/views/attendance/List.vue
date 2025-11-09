<script setup lang="ts">
import AttendanceEditor, {type Form} from '@/components/editor/AttendanceEditor.vue';
import Tag from '@/components/Tag.vue';
import {request} from '@/utils/axios';
import type {Attendance, AttendanceType} from '@/utils/tables';
import {type List} from '@/utils/utils';
import {Plus, Search} from '@element-plus/icons-vue';
import {useUrlSearchParams} from '@vueuse/core';
import dayjs from 'dayjs';
import {reactive, ref, watchEffect} from 'vue';

const params = useUrlSearchParams<Record<string, string>>('history')

const attendances = ref<List<Attendance> | null>(null)
watchEffect(async () => attendances.value = await request('GET', '/attendances', params))

const attendanceTypes = ref<AttendanceType[]>([])
request<AttendanceType[]>('GET', '/attendance-types').then(res => attendanceTypes.value = res)

const isDialogOpen = ref(false)
const attendanceForm = reactive<Form>({
  name: '', location: '', typeId: '', free: false, time: [new Date(), new Date()],
})
</script>

<template>

  <div class="max-w-5xl mx-auto flex flex-col gap-4 p-4">

    <div class="flex justify-between">
      <div class="text-xl">{{$t('attendanceList.title')}}</div>
      <el-button :icon="Plus" @click="isDialogOpen = true">
        {{$t('attendanceList.addButton')}}
      </el-button>
    </div>

    <div class="flex gap-4">
      <el-input v-model="params.name" :prefix-icon="Search" :placeholder="$t('attendanceList.search')">
      </el-input>
      <el-date-picker type="datetimerange" :start-placeholder="$t('attendanceList.startTime')" :end-placeholder="$t('attendanceList.endTime')">
      </el-date-picker>
      <el-button @click="delete params.typeId">{{$t('attendanceList.typeAll')}}</el-button>
      <el-segmented v-model="params.typeId" :options="attendanceTypes" :props="{ value: 'id' }">
        <template #default="{item}">
          {{item.attendanceCount}} {{item.name}}
        </template>
      </el-segmented>
    </div>

    <div v-if="attendances" class="grid grid-cols-[repeat(auto-fill,minmax(250px,1fr))] gap-4">
      <div v-for="a in attendances.data" class="card cursor-pointer flex flex-col gap-2" @click="$router.push(`/attendances/${a.id}`)">
        <div class="flex gap-2 items-center">
          <tag class="shrink-0" :color="a.type.color">{{a.type.name}}</tag>
          <span class="truncate">{{a.name}}</span>
        </div>
        <div class="text-subtle line-clamp-2">{{a.location}}</div>
        <div class="flex gap-2 justify-between text-subtle">
          <div>{{dayjs(a.startTime).format('YYYY-MM-DD hh:mm')}} ~ {{dayjs(a.endTime).format('hh:mm')}}</div>
          <div>{{$t('attendanceList.userCount', { v: a.userCount })}}</div>
        </div>
      </div>
    </div>

    <el-pagination v-if="attendances" layout="prev, pager, next, total" :total="attendances.total"
      :current-page="Number(params.page || 1)" :page-size="Number(params.size || 9)"
      @update:current-page="v => params.page = String(v)" @update:page-size="v => params.size = String(v)"
    ></el-pagination>

  </div>

  <el-dialog v-model="isDialogOpen" class="w-dialog" :title="$t('attendanceList.dialogTitle')">
    <attendance-editor v-model="attendanceForm">
    </attendance-editor>
    <template #footer>
      <el-button type="primary">
        {{$t('confirm')}}
      </el-button>
      <el-button @click="isDialogOpen = false">
        {{$t('cancel')}}
      </el-button>
    </template>
  </el-dialog>

</template>
