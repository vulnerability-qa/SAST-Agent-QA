// CWE-915: Mass Assignment in NestJS controller
import { Body, Patch, Param } from '@nestjs/common';
class UpdateUserDto { [key: string]: unknown; }
class UserService { update(id: string, dto: UpdateUserDto) { return dto; } }
class UserController {
  constructor(private svc: UserService) {}
  @Patch(':id')
  update(@Param('id') id: string, @Body() dto: UpdateUserDto) {
    return this.svc.update(id, dto); // isAdmin settable via body
  }
}
