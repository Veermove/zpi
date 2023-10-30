package zpi.ppea.clap.dtos;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.List;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class RatingsSubmissionResponseDto {
    private List<CriterionDto> criteria;
    private List<AssessorRatingsDto> individualRatings;
    private AssessorRatingsDto initialRating;
    private AssessorRatingsDto finalRating;
}
