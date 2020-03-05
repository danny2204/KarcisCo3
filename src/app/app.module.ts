import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { MainHomeComponent } from './Components/Customer/main-home/main-home.component';
import { HeaderComponent } from './Components/Customer/header/header.component';
import { FooterComponent } from './Components/Customer/footer/footer.component';
import { AdminHomeComponent } from './Components/Admin/admin-home/admin-home.component';
import { RegisterDialogComponent } from './Components/Customer/register-dialog/register-dialog.component';
import { LoginDialogComponent } from './Components/Customer/login-dialog/login-dialog.component';
import { ModalModule } from 'angular-custom-modal';
import { MainPageComponent } from './Components/Customer/main-page/main-page.component';
import { PesawatComponent } from './Components/Customer/child-component/pesawat/pesawat.component';
import { HotelComponent } from './Components/Customer/child-component/hotel/hotel.component';
import { KeretaApiComponent } from './Components/Customer/child-component/kereta-api/kereta-api.component';
import { SewaMobilComponent } from './Components/Customer/child-component/sewa-mobil/sewa-mobil.component';
import { HiburanComponent } from './Components/Customer/child-component/hiburan/hiburan.component';
import { VendorDetailsComponent } from './Components/Customer/child-component/vendor-details/vendor-details.component';
import { MatDialogModule, MatNativeDateModule, MatFormFieldModule, MatInputModule, MatOptionModule, MatSelectModule, MatAccordion, MatExpansionModule, MatTableModule, MatSpinner, MatProgressSpinnerModule } from '@angular/material';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { QuickCardComponent } from './Components/Customer/main-page/components/quick-card/quick-card.component';
import { QuickCardListComponent } from './Components/Customer/main-page/components/quick-card-content/quick-card-list/quick-card-list.component';
import { QuickCardPesawatComponent } from './Components/Customer/main-page/components/quick-card-content/quick-card-list/quick-card-pesawat/quick-card-pesawat.component';
import { GoogleButtonComponent } from './Components/Customer/login-dialog/google-button/google-button.component';
import { FacebookButtonComponent } from './Components/Customer/login-dialog/facebook-button/facebook-button.component';
import { ImageSliderComponent } from './Components/Customer/main-page/components/image-slider/image-slider.component'
import { FormsModule } from '@angular/forms';
import { GraphQLModule } from './graphql.module';
import { HttpClientModule } from '@angular/common/http';
import { QuickCardHotelComponent } from './Components/Customer/main-page/components/quick-card-content/quick-card-list/quick-card-hotel/quick-card-hotel.component';
import { CommonModule } from '@angular/common';
import { MaterialModule } from './material/material.module';
import { PesawatWidgetComponent } from './Components/Customer/child-component/pesawat/pesawat-widget/pesawat-widget.component';
import { QuickCardKeretaComponent } from './Components/Customer/main-page/components/quick-card-content/quick-card-list/quick-card-kereta/quick-card-kereta.component';
import { QuickCardMobilComponent } from './Components/Customer/main-page/components/quick-card-content/quick-card-list/quick-card-mobil/quick-card-mobil.component';
import { HotelDetailComponent } from './Components/Customer/child-component/hotel/hotel-detail/hotel-detail.component';
import { HotelMapComponent } from './Components/Customer/child-component/hotel/hotel-map/hotel-map.component';
import { PopUpImageComponent } from './Components/Customer/child-component/hotel/hotel-detail/pop-up-image/pop-up-image.component';
import { CheckoutPageComponent } from './Components/Customer/child-component/checkout-page/checkout-page.component';
import { AdminMainComponent } from './Components/Admin/admin-main/admin-main.component'
import { HistorySearchComponent } from './Components/Customer/main-page/history-search/history-search.component';
import { PromoPageComponent } from './Components/Customer/promo-page/promo-page.component';
import { PromoDetailPageComponent } from './Components/Customer/promo-page/promo-detail-page/promo-detail-page.component';
import { BlogPageComponent } from './Components/Customer/blog-page/blog-page.component';
import { BlogDetailComponent } from './Components/Customer/blog-page/blog-detail/blog-detail.component';
import { HiburanDetailComponent } from './Components/Customer/child-component/hiburan/hiburan-detail/hiburan-detail.component';
import { HiburanSearchComponent } from './Components/Customer/child-component/hiburan/hiburan-search/hiburan-search.component';
import { UserPageComponent } from './Components/Customer/child-component/user-page/user-page.component';
import { CalendarModule, DateAdapter } from 'angular-calendar';
import { adapterFactory } from 'angular-calendar/date-adapters/date-fns';
import { ChatDetailComponent } from './Components/Customer/chat-page/chat-detail/chat-detail.component';
import { ChatPageComponent } from './Components/Customer/chat-page/chat-page.component';
import { AdminLoginComponent } from './Components/Admin/admin-login/admin-login.component';

@NgModule({
  declarations: [
    AppComponent,
    MainHomeComponent,
    HeaderComponent,
    FooterComponent,
    AdminHomeComponent,
    RegisterDialogComponent,
    LoginDialogComponent,
    MainPageComponent,
    PesawatComponent,
    HotelComponent,
    KeretaApiComponent,
    SewaMobilComponent,
    HiburanComponent,
    VendorDetailsComponent,
    HotelDetailComponent,
    QuickCardComponent,
    QuickCardListComponent,
    QuickCardPesawatComponent,
    QuickCardHotelComponent,
    QuickCardKeretaComponent,
    QuickCardMobilComponent,
    GoogleButtonComponent,
    FacebookButtonComponent,
    ImageSliderComponent,
    PesawatWidgetComponent,
    HotelMapComponent,
    PopUpImageComponent,
    CheckoutPageComponent,
    AdminMainComponent,
    HistorySearchComponent,
    PromoPageComponent,
    PromoDetailPageComponent,
    BlogPageComponent,
    BlogDetailComponent,
    HiburanSearchComponent,
    HiburanDetailComponent,
    UserPageComponent,
    ChatPageComponent,
    ChatDetailComponent,
    AdminLoginComponent
  ],
  entryComponents: [LoginDialogComponent, RegisterDialogComponent, QuickCardPesawatComponent, QuickCardHotelComponent, PesawatWidgetComponent, QuickCardKeretaComponent, QuickCardMobilComponent, PopUpImageComponent, HistorySearchComponent],
  imports: [
    BrowserModule,
    AppRoutingModule,
    ModalModule,
    MatDialogModule,
    BrowserAnimationsModule,
    FormsModule,
    GraphQLModule,
    HttpClientModule,
    CommonModule,
    MaterialModule,
    MatExpansionModule,
    MatTableModule,
    MatProgressSpinnerModule,
    CalendarModule.forRoot({ provide: DateAdapter, useFactory: adapterFactory }),
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
