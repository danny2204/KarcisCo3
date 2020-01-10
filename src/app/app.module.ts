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
import { MatDialogModule, MatNativeDateModule, MatFormFieldModule, MatInputModule } from '@angular/material';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { QuickCardComponent } from './Components/Customer/main-page/components/quick-card/quick-card.component';
import { MatTabsModule } from '@angular/material/tabs';
import { QuickCardListComponent } from './Components/Customer/main-page/components/quick-card-content/quick-card-list/quick-card-list.component';
import { QuickCardPesawatComponent } from './Components/Customer/main-page/components/quick-card-content/quick-card-list/quick-card-pesawat/quick-card-pesawat.component';
import { GoogleButtonComponent } from './Components/Customer/login-dialog/google-button/google-button.component';
import { FacebookButtonComponent } from './Components/Customer/login-dialog/facebook-button/facebook-button.component';
import { ImageSliderComponent } from './Components/Customer/main-page/components/image-slider/image-slider.component'
import { MatDatepickerModule } from '@angular/material/datepicker'
import { FormsModule } from '@angular/forms';
import { GraphQLModule } from './graphql.module';
import { HttpClientModule } from '@angular/common/http';
import { QuickCardHotelComponent } from './Components/Customer/main-page/components/quick-card-content/quick-card-list/quick-card-hotel/quick-card-hotel.component';

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
    QuickCardComponent,
    QuickCardListComponent,
    QuickCardPesawatComponent,
    QuickCardHotelComponent,
    GoogleButtonComponent,
    FacebookButtonComponent,
    ImageSliderComponent,
  ],
  entryComponents: [LoginDialogComponent, RegisterDialogComponent, QuickCardPesawatComponent, QuickCardHotelComponent],
  imports: [
    BrowserModule,
    AppRoutingModule,
    ModalModule,
    MatDialogModule,
    BrowserAnimationsModule,
    MatTabsModule,
    MatDatepickerModule,
    MatNativeDateModule,
    FormsModule,
    MatFormFieldModule,
    MatInputModule,
    GraphQLModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
