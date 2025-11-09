<script setup lang="ts">
import {request} from '@/utils/axios';
import type {Record} from '@/utils/tables';
import {useUrlSearchParams} from '@vueuse/core';
import {reactive, ref, watchEffect} from 'vue';
import {Search} from '@element-plus/icons-vue';
import {shortcuts, type List} from '@/utils/utils';
import RecordEditor, {type Form} from '@/components/editor/RecordEditor.vue';
import {useRouteParams} from '@vueuse/router';

const records = ref<List<Record> | null>(null);

const id = useRouteParams('id')
const params = useUrlSearchParams<globalThis.Record<string, string>>('history');
watchEffect(async () => records.value = await request('GET', `/tasks/${id.value}/records`, params));

const isRecordDialogOpen = ref(false)
const form = reactive<Form>({
  title: '', content: '', minutes: 0,
})
</script>

<template>

  <div class="flex justify-between">
    <div class="text-xl">{{$t('record')}}</div>
    <el-button @click="isRecordDialogOpen = true">
      {{$t('add', { name: $t('record') })}}
    </el-button>
  </div>

  <div class="space-x-4">
    <el-input
      v-model="params.title"
      class="!w-64"
      :prefix-icon="Search"
      :placeholder="$t('search', { v: $t('title') })"
    />
    <el-date-picker type="datetimerange"
      :model-value="params.time"
      @update:model-value="v => params.time = v?.map((e: Date) => e.toISOString())"
      clearable :shortcuts="shortcuts"
      :start-placeholder="$t('select', { v: $t('startTime') })"
      :end-placeholder="$t('select', { v: $t('endTime') })"
    ></el-date-picker>
  </div>

  <!-- Record Cards -->
  <div v-if="records" class="grid grid-cols-[repeat(auto-fill,minmax(300px,1fr))] gap-4">
    <div
      v-for="r in records.data"
      :key="r.id"
      @click="$router.push(`/records/${r.id}`)"
      class="card cursor-pointer flex flex-col gap-2"
    >
      <div class="flex justify-between items-center">
        <div class="font-semibold line-clamp-1">{{r.title}}</div>
        <div class="text-subtle">{{r.createdAt.split('T')[0]}}</div>
      </div>
      <div class="flex items-center gap-2 justify-between">
        <span>{{r.user.name}}</span>
        <span class="text-subtle">
          {{$t('hours')}}: {{Math.floor(r.minutes / 60)}} {{$t('hour')}} {{r.minutes % 60}} {{$t('minutes')}}
        </span>
      </div>
      <div class="text-subtle line-clamp-3">
        {{r.content}}
      </div>
    </div>
  </div>

  <el-pagination v-if="records"
    class="card"
    layout="prev, pager, next, total"
    :total="records.total"
    :current-page="Number(params.page || 1)"
    :page-size="Number(params.size || 9)"
    @update:current-page="v => params.page = String(v)"
    @update:page-size="v => params.size = String(v)"
  />

  <el-dialog v-model="isRecordDialogOpen" :title="$t('taskDetail.recordDialogTitle')">
    <record-editor v-model="form" />
    <template #footer>
        <el-button type="primary">
        {{$t('confirm')}}
      </el-button>
    </template>
  </el-dialog>

</template>

