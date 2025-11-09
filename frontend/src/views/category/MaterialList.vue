<script setup lang="ts">
import {useSessionStore} from '@/stores/session';
import {request} from '@/utils/axios';
import type {Material} from '@/utils/tables';
import {formatDate, type List} from '@/utils/utils';
import {useUrlSearchParams} from '@vueuse/core';
import {ref, watchEffect} from 'vue';
import {useRoute} from 'vue-router';

const route = useRoute()
const materials = ref<List<Material> | null>(null)
const session = useSessionStore()
const params = useUrlSearchParams<Record<string, string>>('history')
watchEffect(async () => materials.value = await request('GET', '/materials', { categoryId: route.params.id, ...params }))
</script>

<template>


    <div class="flex flex-wrap gap-2">
      <el-input class="!w-64" v-model="params.name" :placeholder="$t('search', { v: $t('material') })">
      </el-input>
      <el-select class="!w-32" v-model="params.statusId" :placeholder="$t('select', { v: $t('status') })">
        <el-option :label="$t('all', { v: $t('status') })" :value="0">
        </el-option>
        <el-option v-for="v in session.enums['material_statuses']" :key="v.id" :label="v.name" :value="v.id">
        </el-option>
      </el-select>
    </div>

    <div v-if="materials" class="grid grid-cols-[repeat(auto-fill,minmax(250px,1fr))] gap-4">
      <div v-for="m in materials.data" :key="m.id" @click="$router.push(`/materials/${m.id}`)" class="card space-y-2 cursor-pointer">
        <div class="space-x-2">
          <tag :color="m.status.color">{{m.status.name}}</tag>
          <span>{{m.name}} {{m.quantity}} {{m.unit}}</span>
        </div>
        <div class="text-subtle line-clamp-2">{{m.profile}}</div>
        <div class="text-subtle space-x-2">
          <span>{{$t('count', { k: $t('purchase'), v: m.purchaseCount } )}}</span>
          <span>{{$t('count', { k: $t('task'), v: m.taskCount })}}</span>
        </div>
        <div class="text-subtle">{{$t('updatedAt')}} {{formatDate(m.updatedAt)}}</div>
      </div>
    </div>

    <el-pagination v-if="materials" class="card" layout="prev, pager, next, total" :total="materials.total"
      :current-page="Number(params.page ?? 1)" @update:current-page="v => params.page = String(v)"
      :page-size="Number(params.size ?? 9)" @update:page-size="v => params.size = String(v)"
    ></el-pagination>
</template>
