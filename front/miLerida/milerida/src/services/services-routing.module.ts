import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

const routes: Routes = [
  { path: '', redirectTo: 'home', pathMatch: 'full' },
  { path: 'home', loadChildren: () => import('../home/home.module').then(m => m.HomeModule) },
  //{ path: 'restaurants', loadChildren: () => import('../restaurants/restaurants.module').then(m => m.RestaurantsModule) },
  { path: 'transport', loadChildren: () => import('../transport/transport.module').then(m => m.TransportModule) },
  { path: 'services', loadChildren: () => import('../services/services.module').then(m => m.ServicesModule) },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
