import { buildAuthorizationHeader as buildDefaultHeaders } from "../util";
import { clapConfig } from "./config";
import { RatingsDTO } from "./model/rating";
import { SubmissionDTO, SubmissionDetailsDTO } from "./model/submission";

export class ClapApi {
  private readonly defaultHeaders: HeadersInit;
  private readonly baseUrl = clapConfig.baseUrl;

  constructor(idToken: string) {
    this.defaultHeaders = buildDefaultHeaders(idToken);
    console.log('baseUrl', this.baseUrl)
  }

  async getSubmissions(): Promise<SubmissionDTO[]> {
    const response = await fetch(`${this.baseUrl}/submissions`, {
      headers: this.defaultHeaders
    });
    return await response.json();
  }

  async getSubmissionDetails(submissionId: number): Promise<SubmissionDetailsDTO> {
    const response = await fetch(`${this.baseUrl}/submissions/${submissionId}`, {
      headers: this.defaultHeaders
    });
    return await response.json();
  }

  async getSubmissionRatings(submissionId: number): Promise<RatingsDTO> {
    const response = await fetch(`${this.baseUrl}/submissions/${submissionId}/ratings`, {
      headers: this.defaultHeaders
    });
    return await response.json();
  }
}
