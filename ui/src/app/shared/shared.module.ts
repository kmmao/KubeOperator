import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {CoreModule} from '../core/core.module';
import {ModalAlertComponent} from './common-component/modal-alert/modal-alert.component';
import {K8sPaginationComponent} from './common-component/k8s-pagination/k8s-pagination.component';
import {ZoneStatusPipe} from './pipe/zone-status.pipe';
import {MenuAuthDirective} from './directive/menu-auth.directive';
import {OperateAuthDirective} from './directive/operate-auth.directive';


@NgModule({
    declarations: [ModalAlertComponent, K8sPaginationComponent, ZoneStatusPipe, MenuAuthDirective, OperateAuthDirective
    ],
    exports: [
        ModalAlertComponent,
        K8sPaginationComponent,
        ZoneStatusPipe,
        MenuAuthDirective,
        OperateAuthDirective
    ],
    imports: [
        CommonModule,
        CoreModule,
    ]
})
export class SharedModule {
}
