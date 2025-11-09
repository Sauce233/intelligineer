<script setup lang="ts">
import Tag from '@/components/Tag.vue';
import {request} from '@/utils/axios';
import type {ApplicationStatus, InvoiceStatus, MaterialStatus, ProjectStatus, PurchaseStatus, TaskStatus} from '@/utils/tables';
import type {EChartsOption} from 'echarts';
import {PieChart} from 'echarts/charts';
import {LegendComponent, TitleComponent, TooltipComponent} from 'echarts/components';
import {use} from 'echarts/core';
import {CanvasRenderer} from 'echarts/renderers';
import {computed, ref} from 'vue';
import {LabelLayout} from 'echarts/features';
import VChart from 'vue-echarts';

const colors: Record<string, string> = {
  red: "#ef4444",
  orange: "#f97316",
  amber: "#f59e0b",
  yellow: "#eab308",
  lime: "#84cc16",
  green: "#22c55e",
  emerald: "#10b981",
  teal: "#14b8a6",
  cyan: "#06b6d4",
  sky: "#0ea5e9",
  blue: "#3b82f6",
  indigo: "#6366f1",
  violet: "#8b5cf6",
  purple: "#a855f7",
  fuchsia: "#d946ef",
  pink: "#ec4899",
  rose: "#f43f5e",
  slate: "#64748b",
  gray: "#6b7280",
  zinc: "#71717a",
  neutral: "#737373",
  stone: "#78716c"
};


use([TitleComponent, TooltipComponent, LegendComponent, PieChart, CanvasRenderer, LabelLayout])

const projectStatuses = ref<ProjectStatus[]>([])
request<ProjectStatus[]>('GET', '/project-statuses').then(res => projectStatuses.value = res)
const projectStatusOption = computed<EChartsOption>(() => ({
  title: { text: '專案狀態佔比', left: 'center' },
  tooltip: { trigger: 'item' },
  series: [{
    type: 'pie',
    data: projectStatuses.value.map((s: ProjectStatus) => ({ name: s.name, value: s.projectCount, itemStyle: { color: colors[s.color] } })),
    label: { position: 'inside', formatter: '{b}\n{d}%', color: '#fff', fontSize: 14 },
  }]
}))

const taskStatuses = ref<TaskStatus[]>([])
request<TaskStatus[]>('GET', '/task-statuses').then(res => taskStatuses.value = res)
const taskStatusOption = computed<EChartsOption>(() => ({
  title: { text: '任務狀態佔比', left: 'center' },
  tooltip: { trigger: 'item' },
  series: [{
    type: 'pie',
    data: taskStatuses.value.map((s: TaskStatus) => ({ name: s.name, value: s.taskCount, itemStyle: { color: colors[s.color] } })),
    label: { position: 'inside', formatter: '{b}\n{d}%', color: '#fff', fontSize: 14 },
  }]
}))

const invoiceStatuses = ref<InvoiceStatus[]>([])
request<InvoiceStatus[]>('GET', '/invoice-statuses').then(res => invoiceStatuses.value = res)
const invoiceStatusOption = computed<EChartsOption>(() => ({
  title: { text: '發票狀態佔比', left: 'center' },
  tooltip: { trigger: 'item' },
  series: [{
    type: 'pie',
    data: invoiceStatuses.value.map((s: InvoiceStatus) => ({ name: s.name, value: s.invoiceCount, itemStyle: { color: colors[s.color] } })),
    label: { position: 'inside', formatter: '{b}\n{d}%', color: '#fff', fontSize: 14 },
  }]
}))

const applicationStatuses = ref<ApplicationStatus[]>([])
request<ApplicationStatus[]>('GET', '/application-statuses').then(res => applicationStatuses.value = res)
const applicationStatusOption = computed<EChartsOption>(() => ({
  title: { text: '申請狀態佔比', left: 'center' },
  tooltip: { trigger: 'item' },
  series: [{
    type: 'pie',
    data: applicationStatuses.value.map((s: ApplicationStatus) => ({ name: s.name, value: s.applicationCount, itemStyle: { color: colors[s.color] } })),
    label: { position: 'inside', formatter: '{b}\n{d}%', color: '#fff', fontSize: 14 },
  }]
}))

const purchaseStatuses = ref<PurchaseStatus[]>([])
request<PurchaseStatus[]>('GET', '/purchase-statuses').then(res => purchaseStatuses.value = res)
const purchaseStatusOption = computed<EChartsOption>(() => ({
  title: { text: '採購狀態佔比', left: 'center' },
  tooltip: { trigger: 'item' },
  series: [{
    type: 'pie',
    data: purchaseStatuses.value.map((s: PurchaseStatus) => ({ name: s.name, value: s.purchaseCount, itemStyle: { color: colors[s.color] } })),
    label: { position: 'inside', formatter: '{b}\n{d}%', color: '#fff', fontSize: 14 },
  }]
}))

const materialStatuses = ref<MaterialStatus[]>([])
request<MaterialStatus[]>('GET', '/material-statuses').then(res => materialStatuses.value = res)
const materialStatusOption = computed<EChartsOption>(() => ({
  title: { text: '物料狀態佔比', left: 'center' },
  tooltip: { trigger: 'item' },
  series: [{
    type: 'pie',
    data: materialStatuses.value.map((s: MaterialStatus) => ({ name: s.name, value: s.materialCount, itemStyle: { color: colors[s.color] } })),
    label: { position: 'inside', formatter: '{b}\n{d}%', color: '#fff', fontSize: 14 },
  }]
}))
</script>

<template>
  <div class="grid grid-cols-[repeat(auto-fill,minmax(350px,1fr))] gap-4 p-4">

    <div class="card flex flex-col gap-4">
      <div class="text-lg font-bold">{{$t('dashboard.project.title')}}</div>
      <div class="grow flex gap-4">
        <div class="flex flex-col gap-2 shrink-0">
          <div v-for="ps in projectStatuses" :key="ps.id" class="flex gap-2 items-center">
            <tag :color="ps.color">{{ps.name}}</tag>
            <span class="text-3xl">{{ps.projectCount}}</span>
          </div>
        </div>
        <v-chart :option="projectStatusOption" autoresize>
        </v-chart>
      </div>
    </div>

    <div class="card flex flex-col gap-4">
      <div class="text-lg font-bold">{{$t('dashboard.task.title')}}</div>
      <div class="grow flex gap-4">
        <div class="shrink-0 flex flex-col gap-2">
          <div v-for="ts in taskStatuses" :key="ts.id" class="flex gap-2 items-center">
            <tag :color="ts.color">{{ts.name}}</tag>
            <span class="text-3xl">{{ts.taskCount}}</span>
          </div>
        </div>
        <v-chart :option="taskStatusOption" autoresize>
        </v-chart>
      </div>
    </div>

    <div class="card flex flex-col gap-4">
      <div class="text-lg font-bold">{{$t('dashboard.invoice.title')}}</div>
      <div class="grow flex gap-4">
        <div class="shrink-0 flex flex-col gap-2">
          <div v-for="is in invoiceStatuses" :key="is.id" class="flex gap-2 items-center">
            <span class="text-xl">{{is.invoiceCount}}</span> 張
            <tag :color="is.color">{{is.name}}</tag>
            總價 ¥ <span class="text-xl">{{is.totalAmount}}</span>
          </div>
        </div>
        <v-chart :option="invoiceStatusOption" autoresize>
        </v-chart>
      </div>
    </div>

    <div class="card flex flex-col gap-4">
      <div class="text-lg font-bold">{{$t('dashboard.application.title')}}</div>
      <div class="grow flex gap-4">
        <div class="shrink-0 flex flex-col gap-2">
          <div v-for="as in applicationStatuses" :key="as.id" class="flex gap-2 items-center">
            <span class="text-2xl">{{as.applicationCount}}</span>
            <tag :color="as.color">{{as.name}}</tag>
          </div>
        </div>
        <v-chart :option="applicationStatusOption" autoresize>
        </v-chart>
      </div>
    </div>

    <div class="card flex flex-col gap-4">
      <div class="text-lg font-bold">{{$t('dashboard.purchase.title')}}</div>
      <div class="grow flex gap-4">
        <div class="shrink-0 flex flex-col gap-2">
          <div v-for="ps in purchaseStatuses" :key="ps.id" class="flex gap-2 items-center">
            <span class="text-2xl">{{ps.purchaseCount}}</span>
            <tag :color="ps.color">{{ps.name}}</tag>
          </div>
        </div>
        <v-chart :option="purchaseStatusOption" autoresize>
        </v-chart>
      </div>
    </div>

    <div class="card flex flex-col gap-4">
      <div class="text-lg font-bold">{{$t('dashboard.material.title')}}</div>
      <div class="grow flex gap-4">
        <div class="shrink-0 flex flex-col gap-2">
          <div v-for="ms in materialStatuses" :key="ms.id" class="flex gap-2 items-center">
            <span class="text-2xl">{{ms.materialCount}}</span>
            <tag :color="ms.color">{{ms.name}}</tag>
          </div>
        </div>
        <v-chart :option="materialStatusOption" autoresize>
        </v-chart>
      </div>
    </div>

  </div>
</template>
