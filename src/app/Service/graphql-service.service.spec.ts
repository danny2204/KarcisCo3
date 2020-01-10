import { TestBed } from '@angular/core/testing';

import { GraphqlServiceService } from './graphql-service.service';

describe('GraphqlServiceService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: GraphqlServiceService = TestBed.get(GraphqlServiceService);
    expect(service).toBeTruthy();
  });
});
