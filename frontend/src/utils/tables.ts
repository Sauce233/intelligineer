export interface Client {
  id: number;
  createdAt: string;
  updatedAt: string;
  name: string;
  address: string;
  profile: string;
  // Sub-tables for Client
  contactCount: number;
  contacts: ClientContact[];
  projectCount: number;
  projects: Project[];
}

export interface ClientContact {
  id: number;
  createdAt: string;
  updatedAt: string;
  name: string;
  phone: string;
  email: string;
  profile: string;
  client: Client;
}

export interface Attachment {
  id: number;
  createdAt: string;
  name: string;
  filename: string;
  project: Project;
}

export interface ProjectStatus {
  id: number;
  name: string;
  color: string;
  projectCount: number;
}

export interface Project {
  id: number;
  createdAt: string;
  updatedAt: string;
  name: string;
  client: Client;
  profile: string;
  offer: number;
  budget: number;
  status: ProjectStatus;
  // Sub-tables for Project
  attachmentCount: number;
  attachments: Attachment[];
  invoiceCount: number;
  invoices: Invoice[];
  taskCount: number;
  tasks: Task[];
}

export interface InvoiceStatus {
  id: number;
  name: string;
  color: string;
  invoiceCount: number;
  totalAmount: number;
}

export interface Invoice {
  id: number;
  createdAt: string;
  project: Project;
  name: string;
  number: string;
  profile: string;
  amount: number;
  issueDate: string;
  dueDate: string;
  status: InvoiceStatus;
}

// pending, in_progress, completed
export interface TaskStatus {
  id: number;
  name: string;
  color: string;
  taskCount: number;
}

// start_time and end_time are estimated if the task has not started or finished
export interface Task {
  id: number;
  createdAt: string;
  updatedAt: string;
  name: string;
  profile: string;
  project: Project;
  startTime: string;
  endTime: string;
  status: TaskStatus;
  // Sub-tables for Task
  recordCount: number;
  records: Record[];
  materialCount: number;
  materials: TaskMaterial[];
}

export interface Supplier {
  id: number;
  createdAt: string;
  updatedAt: string;
  name: string;
  address: string;
  profile: string;
  // Sub-tables for Supplier
  contactCount: number;
  contacts: SupplierContact[];
  purchaseCount: number;
  purchases: Purchase[];
}

export interface SupplierContact {
  id: number;
  createdAt: string;
  updatedAt: string;
  name: string;
  phone: string;
  email: string;
  profile: string;
  supplier: Supplier;
}

export interface Category {
  id: number;
  name: string;
  profile: string;
  materialCount: number;
  materials: Material[];
}

export interface MaterialStatus {
  id: number;
  name: string;
  color: string;
  materialCount: number;
}

export interface Material {
  id: number;
  updatedAt: string;
  category: Category;
  status: MaterialStatus;
  name: string;
  quantity: number;
  unit: string;
  profile: string;
  purchaseCount: number;
  taskCount: number;
}

export interface PurchaseStatus {
  id: number;
  name: string;
  color: string;
  purchaseCount: number;
}

export interface Purchase {
  id: number;
  createdAt: string;
  updatedAt: string;
  name: string;
  profile: string;
  supplier: Supplier;
  status: PurchaseStatus;
  // Sub-table for Purchase
  materialCount: number;
  totalPrice: string;
  materials: PurchaseMaterial[];
}

export interface PurchaseMaterial {
  purchase: Purchase;
  material: Material;
  quantity: number;
  price: number;
}

// Quantity represents the amount of material consumed
// If the material's status changes, note it in Notes, e.g., "caused minor damage to the crane"
export interface TaskMaterial {
  task: Task;
  material: Material;
  quantity: number;
  profile: string;
}

export interface User {
  id: number;
  createdAt: string;
  updatedAt: string;
  name: string;
  email: string;
  profile: string;
  // Sub-tables for User
  roleCount: number;
  roles: Role[];
  applicationCount: number;
  recordCount: number;
  applications: Application[];
}

export interface Role {
  id: number;
  name: string;
  color: string;
  userCount: number;
}

export interface UserRole {
  user: User;
  role: Role;
}

export interface Record {
  id: number;
  createdAt: string;
  task: Task;
  user: User;
  title: string;
  content: string;
  minutes: number;
  // Sub-table for Record
  photoCount: number;
  photos: Photo[];
}

export interface Photo {
  id: number;
  createdAt: string;
  record: Record;
}

export interface ApplicationType {
  id: number;
  name: string;
  color: string;
}

// pending, approved, rejected
export interface ApplicationStatus {
  id: number;
  name: string;
  color: string;
  applicationCount: number;
}

export interface Application {
  id: number;
  createdAt: string;
  updatedAt: string;
  applicant: User;
  approver: User;
  title: string;
  content: string;
  reply: string;
  type: ApplicationType;
  status: ApplicationStatus;
}

export interface AttendanceType {
  id: number;
  name: string;
  color: string;
  attendanceCount: number;
}

export interface Attendance {
  id: number;
  createdAt: string;
  updatedAt: string;
  name: string;
  location: string;
  type: AttendanceType;
  startTime: string;
  endTime: string;
  userCount: number;
}

export interface UserAttendanceStatus {
  id: number;
  name: string;
  color: string;
}

export interface UserAttendance {
  user: User;
  attendance: Attendance;
  time: string | null;
  status: UserAttendanceStatus;
}

export interface ProfileType {
  id: number;
  name: string;
  color: string;
}

export interface Profile {
  id: number;
  createdAt: string;
  updatedAt: string;
  name: string;
  type: ProfileType;
  user: User;
  profile: string;
  filename: string;
}
