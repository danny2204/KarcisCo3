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
import { VendorDetailsComponent } from './Components/Customer/child-component/vendor-details/vendor-details.component';
import { HotelDetailComponent } from './Components/Customer/child-component/hotel/hotel-detail/hotel-detail.component';
import { HotelMapComponent } from './Components/Customer/child-component/hotel/hotel-map/hotel-map.component';
import { CheckoutPageComponent } from './Components/Customer/child-component/checkout-page/checkout-page.component';
import { AdminMainComponent } from './Components/Admin/admin-main/admin-main.component';
import { PromoPageComponent } from './Components/Customer/promo-page/promo-page.component';
import { PromoDetailPageComponent } from './Components/Customer/promo-page/promo-detail-page/promo-detail-page.component';
import { BlogPageComponent } from './Components/Customer/blog-page/blog-page.component';
import { BlogDetailComponent } from './Components/Customer/blog-page/blog-detail/blog-detail.component';
import { HiburanSearchComponent } from './Components/Customer/child-component/hiburan/hiburan-search/hiburan-search.component';
import { HiburanDetailComponent } from './Components/Customer/child-component/hiburan/hiburan-detail/hiburan-detail.component';
import { UserPageComponent } from './Components/Customer/child-component/user-page/user-page.component';
import { ChatPageComponent } from './Components/Customer/chat-page/chat-page.component';
import { ChatDetailComponent } from './Components/Customer/chat-page/chat-detail/chat-detail.component';
import { AdminLoginComponent } from './Components/Admin/admin-login/admin-login.component';

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
}, {
  path: 'vendor-detail',
  component: MainHomeComponent,
  children: [{
    path: '',
    component: VendorDetailsComponent,
    outlet: 'mainCustomer'
  }]
}, {
  path: 'hotel-detail',
  component: MainHomeComponent,
  children: [{
    path: '',
    component: HotelDetailComponent,
    outlet: 'mainCustomer'
  }]
}, {
  path: 'hotel-map',
  component: MainHomeComponent,
  children: [{
    path: '',
    component: HotelMapComponent,
    outlet: 'mainCustomer'
  }]
}, {
  path: 'checkout',
  component: MainHomeComponent,
  children: [{
    path: '',
    component: CheckoutPageComponent,
    outlet: 'mainCustomer'
  }]
}, {
  path: 'admin-main',
  component: AdminHomeComponent,
  children: [{
    path: '',
    component: AdminMainComponent,
    outlet: 'mainAdmin'
  }]
}, {
  path: 'event-page',
  component: MainHomeComponent,
  children: [{
    path: '',
    component: HiburanComponent,
    outlet: 'mainCustomer'
  }]
}, {
  path: 'promo-page',
  component: MainHomeComponent,
  children: [{
    path: '',
    component: PromoPageComponent,
    outlet: "mainCustomer"
  }]
}, {
  path: 'promo-detail/:id',
  component: MainHomeComponent,
  children: [{
    path: '',
    component: PromoDetailPageComponent,
    outlet: 'mainCustomer'
  }]
}, {
  path: 'blog-page',
  component: MainHomeComponent,
  children: [{
    path: '',
    component: BlogPageComponent,
    outlet: 'mainCustomer'
  }]
}, {
  path: 'blog-detail',
  component: MainHomeComponent,
  children: [{
    path: '',
    component: BlogDetailComponent,
    outlet: 'mainCustomer'
  }]
}, {
  path: 'blog-detail/:id',
  component: MainHomeComponent,
  children: [{
    path: '',
    component: BlogDetailComponent,
    outlet: 'mainCustomer'
  }]
}, {
  path: 'hiburan-search',
  component: MainHomeComponent,
  children: [{
    path: '',
    component: HiburanSearchComponent,
    outlet: 'mainCustomer'
  }]
}, {
  path: 'hiburan-detail',
  component: MainHomeComponent,
  children: [{
    path: '',
    component: HiburanDetailComponent,
    outlet: 'mainCustomer'
  }]
}, {
  path: 'user-page',
  component: MainHomeComponent,
  children: [{
    path: '',
    component: UserPageComponent,
    outlet: 'mainCustomer'
  }]
}, {
  path: 'chat-page',
  component: MainHomeComponent,
  children: [{
    path: '',
    component: ChatPageComponent,
    outlet: 'mainCustomer'
  }]
}, {
  path: 'chat-detail',
  component: MainHomeComponent,
  children: [{
    path: '',
    component: ChatDetailComponent,
    outlet: 'mainCustomer'
  }]
}, {
  path: 'admin-login',
  component: MainHomeComponent,
  children: [{
    path: '',
    component: AdminLoginComponent,
    outlet: "mainCustomer"
  }]
}];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
