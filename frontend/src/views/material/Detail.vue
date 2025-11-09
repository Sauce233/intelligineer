<script setup lang="ts">
import Tag from '@/components/Tag.vue';
import {request} from '@/utils/axios';
import type {Material, Purchase, Task} from '@/utils/tables';
import {formatDate} from '@/utils/utils';
import {Back, Edit, Plus} from '@element-plus/icons-vue';
import {ref, watchEffect} from 'vue';
import {useRoute} from 'vue-router';

const route = useRoute()

const material = ref<Material | null>(null)

request<Material>('GET', `/materials/${route.params.id}`).then(res => material.value = res)

const purchases = ref<{
  data: Purchase[],
  total: number,
} | null>(null)

watchEffect(async () => purchases.value = await request('GET', '/purchases', { materialId: route.params.id }))

const tasks = ref<{
  data: Task[],
  total: number
} | null>(null)

watchEffect(async () => tasks.value = await request('GET', '/tasks', { materialId: route.params.id }))
</script>

<template>
  <div class="max-w-5xl mx-auto flex flex-col gap-4 p-4">
    <div v-if="material" class="card flex items-center gap-2 flex-wrap">
      <el-button @click="$router.back" :icon="Back">{{$t('back')}}</el-button>
      <div>{{$t('material')}} {{material.name}} {{material.quantity}} {{material.unit}}</div>
      <el-button :icon="Edit" class="ms-auto">{{$t('edit', { v: $t('material') })}}</el-button>
    </div>
    <div v-if="material" class="rounded-2xl p-4 bg-subtle">{{material.profile}}</div>
    <div class="flex justify-between">
      <div class="font-bold text-2xl">{{$t('purchase')}}</div>
      <el-button :icon="Plus">{{$t('add', { v: $t('purchase') })}}</el-button>
    </div>
    <div v-if="purchases" class="grid grid-cols-[repeat(auto-fill,minmax(250px,1fr))] gap-4">
      <div v-for="p in purchases.data" :key="p.id" @click="$router.push(`/purchases/${p.id}`)" class="card cursor-pointer space-y-2">
        <div class="space-x-2">
          <tag :color="p.status.color">{{p.status.name}}</tag>
          <span>{{p.name}}</span>
        </div>
        <div class="text-sm-subtle">{{p.supplier.name}}</div>
        <div class="text-sm-subtle">{{p.profile}}</div>
        <div class="text-sm-subtle">{{$t('materialCount')}} {{p.materialCount}} {{$t('totalPrice')}} ¥ {{p.totalPrice}}</div>
      </div>
    </div>
    <div class="flex justify-between">
      <div class="font-bold text-2xl">{{$t('task')}}</div>
      <el-button :icon="Plus">{{$t('add', { v: $t('task') })}}</el-button>
    </div>
    <div v-if="tasks" class="grid grid-cols-[repeat(auto-fill,minmax(250px,1fr))] gap-4">
      <div v-for="t in tasks.data" :key="t.id" @click="$router.push(`/tasks/${t.id}`)" class="card cursor-pointer space-y-2">
        <div class="space-x-2">
          <tag :color="t.status.color">{{t.status.name}}</tag>
          <span>{{t.name}}</span>
        </div>
        <div class="text-sm-subtle">{{t.profile}}</div>
        <div class="text-sm-subtle">
          {{$t('recordCount')}} {{t.recordCount}} {{$t('materialCount')}} {{t.materialCount}}
        </div>
        <div class="text-sm-subtle">
          {{$t('startTime')}} {{formatDate(t.startTime)}} {{$t('endTime')}} {{formatDate(t.endTime)}}
        </div>
      </div>
    </div>
  </div>
</template>
