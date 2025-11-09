<script setup lang="ts">
import Tag from '@/components/Tag.vue';
import { request } from '@/utils/axios';
import type { Purchase, PurchaseStatus } from '@/utils/tables';
import { Search } from '@element-plus/icons-vue';
import { useUrlSearchParams } from '@vueuse/core';
import { ref, watchEffect } from 'vue';

const purchases = ref<{
  data: Purchase[],
  total: number
} | null>(null)

const params = useUrlSearchParams<Record<string, string>>('history')

// 监听参数变化自动加载采购列表
watchEffect(async () => purchases.value = await request('GET', '/purchases', params))


const purchaseStatuses = ref<PurchaseStatus[]>([])
request<PurchaseStatus[]>('GET', '/purchase-statuses').then(res => purchaseStatuses.value = res)

</script>

<template>

  <div class="p-4 flex flex-col gap-4">

    <div class="text-xl">
      {{$t('purchase')}}
    </div>

    <!-- 搜索与筛选 -->
    <div class="flex gap-4">
      <el-input
        v-model="params.name"
        class="!w-64"
        :prefix-icon="Search"
        :placeholder="$t('search', { v: $t('purchase') })"
      ></el-input>
      <el-select v-model="params.sort" class="!w-32" :placeholder="$t('select', { v: $t('sort') })">
        <el-option v-for="v, k in { created_at: 'createdAt', total_price: 'totalPrice' }" :key="k"
          :label="$t(v)" :value="k"
        ></el-option>
      </el-select>
      <el-select v-model="params.order" class="!w-32" :placeholder="$t('select', { v: $t('order') })">
        <el-option v-for="v in ['asc', 'desc']" :key="v" :label="$t(v)" :value="v">
        </el-option>
      </el-select>
    </div>

    <div class="flex gap-4">
      <el-button @click="delete params.statusId">{{$t('purchaseList.statusAll')}}</el-button>
      <el-segmented v-model="params.statusId" :options="purchaseStatuses" :props="{ value: 'id' }">
        <template #default="{item}">
          {{item.name}} {{item.purchaseCount}}
        </template>
      </el-segmented>
    </div>

    <div v-if="purchases" class="grid grid-cols-[repeat(auto-fill,minmax(250px,1fr))] gap-4">
      <div
        v-for="p in purchases.data"
        :key="p.id"
        @click="$router.push(`/purchases/${p.id}`)"
        class="card cursor-pointer flex flex-col gap-2"
      >
        <div class="space-x-2 line-clamp-1">
          <Tag :color="p.status.color">
            {{ p.status.name }}
          </Tag>
          <span class="text-lg">{{ p.name }}</span>
        </div>

        <div class="text-subtle">{{ p.supplier.name }}</div>

        <div class="line-clamp-2 text-gray-500" v-if="p.profile">
          {{ p.profile }}
        </div>

        <div class="mt-auto flex justify-between text-subtle">
          <div>{{$t('count', { v: p.materialCount, k: $t('material') })}}</div>
          <div>{{$t('totalPrice')}} ¥ {{p.totalPrice}}</div>
          <div>{{ new Date(p.createdAt).toLocaleDateString() }}</div>
        </div>
      </div>
    </div>

      <!-- 分页 -->
    <el-pagination v-if="purchases"
      class="card"
      layout="prev, pager, next, total"
      :total="purchases.total"
      :current-page="Number(params.page || 1)"
      :page-size="Number(params.size || 9)"
      @update:current-page="v => params.page = String(v)"
      @update:page-size="v => params.size = String(v)"
    ></el-pagination>
  </div>

</template>
