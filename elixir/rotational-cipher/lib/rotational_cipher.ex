defmodule RotationalCipher do
  @letters_count 26
  @upper_base ?A
  @upper_range ?A..?Z
  @lower_base ?a
  @lower_range ?a..?z
  @doc """
  Given a plaintext and amount to shift by, return a rotated string.

  Example:
  iex> RotationalCipher.rotate("Attack at dawn", 13)
  "Nggnpx ng qnja"
  """
  @spec rotate(text :: String.t(), shift :: integer) :: String.t()
  def rotate(text, shift) do
    text
    |> to_charlist()
    |> Enum.map(&shift_char(&1, shift))
    |> to_string()
  end

  defp shift_char(char, shift) do
    cond do
      char in @upper_range -> rem(char - @upper_base + shift, @letters_count) + @upper_base
      char in @lower_range -> rem(char - @lower_base + shift, @letters_count) + @lower_base
      true -> char
    end
  end
end
