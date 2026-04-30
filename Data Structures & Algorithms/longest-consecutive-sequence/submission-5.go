type Sequence struct {
	Start int
	End   int
	Size  int
	Next  *Sequence
}

func mergeSequences(sequenceA, sequenceB *Sequence) *Sequence {
	sequenceA.End = sequenceB.End
	sequenceA.Next = sequenceB.Next
	sequenceA.Size = sequenceA.Size + sequenceB.Size
	return sequenceA
}

func addToSequences(sequence *Sequence, element int) (int, *Sequence) {
	if element >= sequence.Start && element <= sequence.End {
		return sequence.Size, sequence
	}

	if element == sequence.Start-1 {
		sequence.Start = element
		sequence.Size++
		return sequence.Size, sequence
	}

	if element == sequence.End+1 {
		sequence.End = element
		sequence.Size++

		if sequence.Next != nil && sequence.End+1 == sequence.Next.Start {
			sequence = mergeSequences(sequence, sequence.Next)
		}

		return sequence.Size, sequence
	}

	if element < sequence.Start-1 {
		newSequence := &Sequence{Start: element, End: element, Size: 1, Next: sequence}

		if newSequence.End+1 == sequence.Start {
			newSequence = mergeSequences(newSequence, sequence)
		}

		return newSequence.Size, newSequence
	}

	if sequence.Next == nil {
		sequence.Next = &Sequence{Start: element, End: element, Size: 1}
		return 1, sequence
	}

	if element < sequence.Next.Start-1 {
		newSequence := &Sequence{Start: element, End: element, Size: 1, Next: sequence.Next}

		if newSequence.End+1 == newSequence.Next.Start {
			newSequence = mergeSequences(newSequence, newSequence.Next)
		}

		sequence.Next = newSequence
		return newSequence.Size, sequence
	}

	currSize, newNext := addToSequences(sequence.Next, element)
	sequence.Next = newNext // IMPORTANT

	if sequence.End+1 == sequence.Next.Start {
		sequence = mergeSequences(sequence, sequence.Next)
		return sequence.Size, sequence
	}

	return currSize, sequence
}

func longestConsecutive(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	maxSequenceSize := 1
	currSequenceSize := 1
	sequence := &Sequence{Start: nums[0], End: nums[0], Size: 1}

	for i := 1; i < len(nums); i++ {
		temp := nums[i]
		currSequenceSize, sequence = addToSequences(sequence, temp)
		if currSequenceSize > maxSequenceSize {
			maxSequenceSize = currSequenceSize
		}
	}

	return maxSequenceSize

}