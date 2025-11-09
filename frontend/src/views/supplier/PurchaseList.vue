<script setup lang="ts">
import PurchaseEditor, {type Form} from '@/components/editor/PurchaseEditor.vue';
import Tag from '@/components/Tag.vue';
import { request } from '@/utils/axios';
import type { Purchase } from '@/utils/tables';
import type {List} from '@/utils/utils';
import { Plus, Search } from '@element-plus/icons-vue';
import { useUrlSearchParams } from '@vueuse/core';
import {useRouteParams} from '@vueuse/router';
import {reactive} from 'vue';
import { ref, watchEffect } from 'vue';

const id = useRouteParams('id')
const purchases = ref<List<Purchase> | null>(null)

const params = useUrlSearchParams<Record<string, string>>('history')

watchEffect(async () => purchases.value = await request('GET', `/suppliers/${id.value}/purchases`, params))

const isDialogOpen = ref(false)
const form = reactive<Form>({
  name: '', profile: '', statusId: 0,
})

</script>

<template>

  <div class="flex justify-between">
    <div class="text-xl">{{$t('purchase')}}</div>
    <el-button :icon="Plus" @click="isDialogOpen = true">
      {{$t('supplier.purchase.addButton')}}
    </el-button>
  </div>

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

      <div class="line-clamp-2 text-subtle" v-if="p.profile">
        {{ p.profile }}
      </div>

      <div class="mt-auto flex justify-between text-subtle">
        <div>{{$t('count', { v: p.materialCount, k: $t('material') })}}</div>
        <div>{{$t('totalPrice')}} ¥ {{p.totalPrice}}</div>
        <div>{{ new Date(p.createdAt).toLocaleDateString() }}</div>
      </div>
    </div>
  </div>

  <el-pagination v-if="purchases"
    class="card"
    layout="prev, pager, next, total"
    :total="purchases.total"
    :current-page="Number(params.page || 1)"
    :page-size="Number(params.size || 9)"
    @update:current-page="v => params.page = String(v)"
    @update:page-size="v => params.size = String(v)"
  ></el-pagination>

  <el-dialog v-model="isDialogOpen" :title="$t('supplier.purchase.dialog.title')">
    <purchase-editor v-model="form">
    </purchase-editor>
    <template #footer>
      <el-button
        type="primary"
        @click="request('POST', `/suppliers/${id}/purchases`, form).then(id => id && $router.push(`/purchases/${id}`))"
      >
        {{$t('confirm')}}
      </el-button>
    </template>
  </el-dialog>

</template>
