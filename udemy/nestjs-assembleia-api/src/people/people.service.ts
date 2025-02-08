import { Injectable } from '@nestjs/common';
import { Person } from './person';

@Injectable()
export class PeopleService {
  private readonly people: Person[] = [];

  list(): Person[] {
    return this.people;
  }

  save(person: Person): void {
    this.people.push(person);
  }
}
