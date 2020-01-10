import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { MainHomeComponent } from './Components/Customer/main-home/main-home.component';
import { AdminHomeComponent } from './Components/Admin/admin-home/admin-home.component';
import { MainPageComponent } from './Components/Customer/main-page/main-page.component';
import { PesawatComponent } from './Components/Customer/child-component/pesawat/pesawat.component';
import { HiburanComponent } from './Components/Customer/child-component/hiburan/hiburan.component';
import { SewaMobilComponent } from './Components/Customer/child-component/sewa-mobil/sewa-mobil.component';
import { HotelComponent } from './Components/Customer/child-component/hotel/hotel.component';
import { KeretaApiComponent } from './Components/Customer/child-component/kereta-api/kereta-api.component';

const routes: Routes = [{
  path: '',
  component: MainHomeComponent,
  children: [{
    path: '',
    component: MainPageComponent,
    outlet: "mainCustomer"
  }]
}, {
  path: 'admin',
  component: AdminHomeComponent,
  children: [{
    path: 'admin',
    component: AdminHomeComponent,
    outlet: "mainAdmin"
  }]
}, {
  path: 'pesawat',
  component: MainHomeComponent,
  children: [{
    path: '',
    component: PesawatComponent,
    outlet: "mainCustomer"
  }]
}, {
  path: 'hiburan',
  component: MainHomeComponent,
  children: [{
    path: '',
    component: HiburanComponent,
    outlet: "mainCustomer"
  }]
}, {
  path: 'sewa-mobil',
  component: MainHomeComponent,
  children: [{
    path: '',
    component: SewaMobilComponent,
    outlet: "mainCustomer"
  }]
}, {
  path: 'hotel',
  component: MainHomeComponent,
  children: [{
    path: '',
    component: HotelComponent,
    outlet: "mainCustomer"
  }]
}, {
  path: 'kereta-api',
  component: MainHomeComponent,
  children: [{
    path: '',
    component: KeretaApiComponent,
    outlet: "mainCustomer"
  }]
}];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
