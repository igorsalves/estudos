import { Injectable, NotFoundException } from '@nestjs/common';
import { Person } from './person';

@Injectable()
export class PeopleService {
  private readonly people: Person[] = [];

  list(): Person[] {
    return this.people;
  }

  findById(id: number): Person {
    const person = this.people.find((person) => person.id === id);

    if (!person) {
      throw new NotFoundException();
    }

    return person;
  }

  save(person: Person): void {
    this.people.push(person);
  }
}
